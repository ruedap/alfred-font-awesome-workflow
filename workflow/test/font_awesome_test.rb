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

  describe '.select!' do
    before { @icons = FontAwesome.icons }
    after  { @icons = nil }

    describe 'with `apple`' do
      before do
        @queries = %w(apple)
        FontAwesome.select!(@icons, @queries)
      end
      after { @queries = nil }

      it { @icons.must_equal @queries }
      it { @icons.count.must_equal 1 }
    end

    describe 'with `left arr`' do
      before do
        @queries = %w(left arr)
        FontAwesome.select!(@icons, @queries)
      end
      after { @queries = nil }

      it { @icons.must_equal %w(arrow-left circle-arrow-left long-arrow-left) }
      it { @icons.count.must_equal 3 }
    end

    describe 'with `arr left` (reverse)' do
      before do
        @queries = %w(arr left)
        FontAwesome.select!(@icons, @queries)
      end
      after { @queries = nil }

      it { @icons.must_equal %w(arrow-left circle-arrow-left long-arrow-left) }
      it { @icons.count.must_equal 3 }
    end

    describe 'with `ruedap` (does not match)' do
      before do
        @queries = %w(ruedap)
        FontAwesome.select!(@icons, @queries)
      end
      after { @queries = nil }

      it { @icons.must_equal %w() }
      it { @icons.must_be_empty }
    end

    describe 'with unknown arguments' do
      before do
        @queries = %w()
        FontAwesome.select!(@icons, @queries)
      end
      after { @queries = nil }

      it { @icons.must_equal FontAwesome.icons }
      it { @icons.count.must_equal 378 }
    end
  end
end
