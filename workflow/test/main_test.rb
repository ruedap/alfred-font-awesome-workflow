require File.expand_path('test_helper', File.dirname(__FILE__))

class MainTest < Minitest::Test
  def test_require_bundler_setup
    assert_equal false, require('bundle/bundler/setup')
  end
end
