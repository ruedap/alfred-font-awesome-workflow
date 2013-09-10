require File.expand_path('test_helper', File.dirname(__FILE__))

describe FontAwesome do
  it 'does not cause an error' do
    require('bundle/bundler/setup').must_equal false
  end

  describe '.icons' do
    before { @icons = FontAwesome.icons }

    it { @icons.size.must_equal 378 }
    it { @icons.first.must_equal 'adjust' }
    it { @icons.last.must_equal 'zoom-out' }

    it 'includes these icons' do
      Fixtures.icons.each { |icon| @icons.must_include icon }
    end

    it 'includes these icons (reverse)' do
      @icons.each { |icon| Fixtures.icons.must_include icon }
    end

    it 'does not includes these icons' do
      icons = %w(icon awesome)
      icons.each { |icon| @icons.wont_include icon }
    end
  end

  describe '.select!' do
    before { @icons = FontAwesome.icons }

    describe 'with `apple`' do
      before do
        @queries = %w(apple)
        FontAwesome.select!(@icons, @queries)
      end

      it { @icons.must_equal @queries }
      it { @icons.size.must_equal 1 }
    end

    describe 'with `left arr`' do
      before do
        @queries = %w(left arr)
        FontAwesome.select!(@icons, @queries)
      end

      it { @icons.must_equal %w(arrow-left circle-arrow-left long-arrow-left) }
      it { @icons.size.must_equal 3 }
    end

    describe 'with `arr left` (reverse)' do
      before do
        @queries = %w(arr left)
        FontAwesome.select!(@icons, @queries)
      end

      it { @icons.must_equal %w(arrow-left circle-arrow-left long-arrow-left) }
      it { @icons.size.must_equal 3 }
    end

    describe 'with `icon` (does not match)' do
      before do
        @queries = %w(icon)
        FontAwesome.select!(@icons, @queries)
      end

      it { @icons.must_equal %w() }
      it { @icons.must_be_empty }
    end

    describe 'with unknown arguments' do
      before do
        @queries = %w()
        FontAwesome.select!(@icons, @queries)
      end

      it { @icons.must_equal FontAwesome.icons }
      it { @icons.size.must_equal 378 }
    end
  end

  describe '.item_hash' do
    before do
      @icon = 'apple'
      @item_hash = FontAwesome.item_hash(@icon)
    end

    it { @item_hash[:uid].must_equal '' }
    it { @item_hash[:title].must_equal @icon }
    it { @item_hash[:subtitle].must_equal "Copy to clipboard: icon-#{@icon}" }
    it { @item_hash[:arg].must_equal @icon }
    it { @item_hash[:icon][:type].must_equal 'default' }
    it { @item_hash[:icon][:name].must_equal "icon-#{@icon}.png" }
    it { @item_hash[:valid].must_equal 'yes' }
    it { @item_hash.size.must_equal 6 }
  end
end
