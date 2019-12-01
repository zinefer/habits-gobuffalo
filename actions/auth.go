package actions

import (
	"fmt"
	"os"
	"time"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/nulls"
	"github.com/markbates/going/defaults"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/facebook"
	"github.com/markbates/goth/providers/github"
	"github.com/markbates/goth/providers/google"
	"github.com/pkg/errors"

	petname "github.com/dustinkirkland/golang-petname"

	"habits/models"
)

func init() {
	gothic.Store = App().SessionStore

	goth.UseProviders(
		google.New(os.Getenv("GOOGLE_KEY"), os.Getenv("GOOGLE_SECRET"), fmt.Sprintf("%s%s", App().Host, "/auth/google/callback")),
		github.New(os.Getenv("GITHUB_KEY"), os.Getenv("GITHUB_SECRET"), fmt.Sprintf("%s%s", App().Host, "/auth/github/callback")),
		facebook.New(os.Getenv("FACEBOOK_KEY"), os.Getenv("FACEBOOK_SECRET"), fmt.Sprintf("%s%s", App().Host, "/auth/facebook/callback")),
	)
}

// AuthCallback is the return oAuth call
func AuthCallback(c buffalo.Context) error {
	gu, err := gothic.CompleteUserAuth(c.Response(), c.Request())
	if err != nil {
		return c.Error(401, err)
	}
	tx := c.Value("tx").(*pop.Connection)
	q := tx.Where("provider = ? and provider_id = ?", gu.Provider, gu.UserID)
	exists, err := q.Exists("users")
	if err != nil {
		return errors.WithStack(err)
	}
	u := &models.User{}
	if exists {
		if err = q.First(u); err != nil {
			return errors.WithStack(err)
		}
	}

	u.Nickname = gu.NickName
	u.Name = defaults.String(gu.Name, gu.NickName)
	u.Provider = gu.Provider
	u.ProviderID = gu.UserID
	u.Email = nulls.NewString(gu.Email)

	if err = preventNicknameCollisions(tx, u, 0); err != nil {
		return errors.WithStack(err)
	}

	if err = tx.Save(u); err != nil {
		return errors.WithStack(err)
	}
	c.Session().Set("current_user_id", u.ID)
	if err = c.Session().Save(); err != nil {
		return errors.WithStack(err)
	}
	c.Flash().Add("success", "You have been logged in")
	return c.Redirect(302, "/")
}

func preventNicknameCollisions(tx *pop.Connection, u *models.User, tries int) error {
	unique := u.HasUniqueNickname(tx)
	if !unique || u.Nickname == "" {
		u.Nickname = petname.Generate(min(tries/5, 2), "-")
		return preventNicknameCollisions(tx, u, tries+1)
	}
	return nil
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

// AuthDestroy kills sessions
func AuthDestroy(c buffalo.Context) error {
	c.Session().Clear()
	c.Flash().Add("success", "You have been logged out")
	return c.Redirect(302, "/")
}

// SetCurrentUser for context/cookies/sessions
func SetCurrentUser(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if uid := c.Session().Get("current_user_id"); uid != nil {
			u := &models.User{}
			tx := c.Value("tx").(*pop.Connection)
			if err := tx.Find(u, uid); err != nil {
				return errors.WithStack(err)
			}

			c.Set("current_user", u)
			c.Cookies().Set("current_user", u.Name, 30*24*time.Hour)
		}
		return next(c)
	}
}

// Authorize ensures the user is signed in
func Authorize(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if uid := c.Session().Get("current_user_id"); uid == nil {
			c.Flash().Add("danger", "You must be authorized to see that page")
			return c.Redirect(302, "/")
		}
		return next(c)
	}
}
