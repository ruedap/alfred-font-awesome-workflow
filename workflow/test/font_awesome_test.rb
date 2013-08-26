require File.expand_path('test_helper', File.dirname(__FILE__))

describe FontAwesome do
  it 'does not cause an error' do
    require('bundle/bundler/setup').must_equal false
  end

  describe '.icons' do
    before { @icons = FontAwesome.icons }
    after  { @icons = nil }

    it { @icons.count.must_equal 378 }

    it { @icons.first.must_equal 'adjust' }

    it { @icons.last.must_equal 'zoom-out' }

    it 'includes these strings' do
      strings = %w(user repeat copy code-fork)
      strings.each { |str| @icons.must_include str }
    end

    it 'does not includes these strings' do
      strings = %w(icon)
      strings.each { |str| @icons.wont_include str }
    end
  end
end
