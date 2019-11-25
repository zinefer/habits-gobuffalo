import { shallowMount } from '@vue/test-utils'

import HomePage from "../../pages/home.vue";
import NavigationComponent from "../../components/navigation.vue";

let wrapper = null

beforeEach(() => {
  wrapper = shallowMount(HomePage)
})

afterEach(() => {
  wrapper.destroy()
})

describe('Home', () => {
  it('renders navigation', () => {
    expect(wrapper.contains(NavigationComponent)).toBe(true)
  })
})