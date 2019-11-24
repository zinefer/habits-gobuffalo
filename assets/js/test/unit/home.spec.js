import { shallowMount } from '@vue/test-utils'
import HomeComponent from "../../components/home.vue";

let wrapper = null

beforeEach(() => {
  wrapper = shallowMount(HomeComponent)
})

afterEach(() => {
  wrapper.destroy()
})

describe('Home', () => {
  it('renders Welcome', () => {
    expect(wrapper.find('.page-header').text()).toContain('Welcome')
  })
})