require File.expand_path('test_helper', File.dirname(__FILE__))

describe FontAwesome do
  it 'does not cause an error' do
    assert_equal false, require('bundle/bundler/setup')
  end

  describe '.icons' do
    before do
      @icons = FontAwesome.icons
    end

    it 'is 378 pieces' do
      @icons.count.must_equal 378
    end

    it 'includes these strings' do
      strings = %w(user repeat copy code-fork)
      strings.each do |str|
        @icons.include?(str).must_equal true
      end
    end

    it 'does not includes these strings' do
      strings = %w(icon)
      strings.each do |str|
        @icons.include?(str).must_equal false
      end
    end

    it 'is "adjust" is the first element' do
      @icons.first.must_equal 'adjust'
    end

    it 'is "zoom-out" is the last element' do
      @icons.last.must_equal 'zoom-out'
    end
  end
end
