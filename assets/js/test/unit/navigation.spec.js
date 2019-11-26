import { shallowMount, createLocalVue } from "@vue/test-utils"
import Buefy from 'buefy'
import { BNavbarItem } from 'buefy/dist/components/navbar'

import NavigationComponent from "../../components/navigation"

const localVue = createLocalVue()
let wrapper = null

localVue.use(Buefy)

beforeEach(() => {
  wrapper = shallowMount(NavigationComponent, { localVue })
})

afterEach(() => {
  wrapper.destroy()
})

describe('Navigation', () => {
  it('renders our name', () => {
    expect(wrapper.find(BNavbarItem).text()).toContain('Habits')
  })
})