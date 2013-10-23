require File.expand_path('test_helper', File.dirname(__FILE__))

describe FontAwesome do
  it 'does not cause an error' do
    require('bundle/bundler/setup').must_equal false
  end

  describe '#icons' do
    before { @icons = FontAwesome.new.icons }

    it { @icons.size.must_equal 409 }
    it { @icons.first.must_equal 'adjust' }
    it { @icons.last.must_equal 'youtube-square' }

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

  describe '#select!' do
    describe 'with `dot-circle-o`' do
      before do
        @queries = %w(dot-circle-o)
        @icons = FontAwesome.new.select!(@queries)
      end

      it { @icons.must_equal @queries }
      it { @icons.size.must_equal 1 }
    end

    describe 'with `left arr`' do
      before do
        queries = %w(left arr)
        @icons = FontAwesome.new.select!(queries)
      end

      it { @icons.must_equal %w(arrow-circle-left arrow-circle-o-left arrow-left long-arrow-left) }
      it { @icons.size.must_equal 4 }
    end

    describe 'with `arr left` (reverse)' do
      before do
        queries = %w(arr left)
        @icons = FontAwesome.new.select!(queries)
      end

      it { @icons.must_equal %w(arrow-circle-left arrow-circle-o-left arrow-left long-arrow-left) }
      it { @icons.size.must_equal 4 }
    end

    describe 'with `icon` (does not match)' do
      before do
        queries = %w(icon)
        @icons = FontAwesome.new.select!(queries)
      end

      it { @icons.must_equal %w() }
      it { @icons.must_be_empty }
    end

    describe 'with unknown arguments' do
      before do
        queries = %w()
        @icons = FontAwesome.new.select!(queries)
      end

      it { @icons.must_equal FontAwesome.new.icons }
      it { @icons.size.must_equal 409 }
    end
  end

  describe '#item_hash' do
    before do
      @icon = 'apple'
      @item_hash = FontAwesome.new.item_hash(@icon)
    end

    it { @item_hash[:uid].must_equal '' }
    it { @item_hash[:title].must_equal @icon }
    it { @item_hash[:subtitle].must_equal "Copy to clipboard: fa-#{@icon}" }
    it { @item_hash[:arg].must_equal @icon }
    it { @item_hash[:icon][:type].must_equal 'default' }
    it { @item_hash[:icon][:name].must_equal "./icons/fa-#{@icon}.png" }
    it { @item_hash[:valid].must_equal 'yes' }
    it { @item_hash.size.must_equal 6 }
  end

  describe '#add_items' do
    before do
      feedback = Alfred::Core.new.feedback
      icons = %w(beer cloud apple)
      @feedback = FontAwesome.new.add_items(feedback, icons)
    end

    it { @feedback.items.first.title.must_equal 'beer' }
    it { @feedback.items.last.title.must_equal 'apple' }
  end

  describe '#to_alfred' do
    before do
      alfred = Alfred::Core.new
      query = 'bookmark'
      @xml = FontAwesome.new(query).to_alfred(alfred)
      @result_xml = "<items><item uid='' valid='yes'><title>bookmark</title><arg>bookmark</arg><subtitle>Copy to clipboard: fa-bookmark</subtitle><icon>./icons/fa-bookmark.png</icon></item><item uid='' valid='yes'><title>bookmark-o</title><arg>bookmark-o</arg><subtitle>Copy to clipboard: fa-bookmark-o</subtitle><icon>./icons/fa-bookmark-o.png</icon></item></items>"
    end

    it { @xml.must_equal @result_xml }
  end
end
