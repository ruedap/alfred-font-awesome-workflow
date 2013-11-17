# encoding: utf-8

require File.expand_path('test_helper', File.dirname(__FILE__))

describe FontAwesome do
  it 'does not cause an error' do
    require('bundle/bundler/setup').must_equal false
  end

  describe '.to_character_reference' do
    it { FontAwesome.to_character_reference('f000').must_equal '' }
    it { FontAwesome.to_character_reference('f17b').must_equal '' }

    it { FontAwesome.to_character_reference('f001').wont_equal '' }
  end

  describe '#icons' do
    before { @icons = FontAwesome.new.icons }

    it { @icons.size.must_equal 409 }
    it { @icons.first.id.must_equal 'adjust' }
    it { @icons.last.id.must_equal 'youtube-square' }

    it 'includes these icons' do
      icon_ids = @icons.map { |icon| icon.id }
      Fixtures.icon_ids.each { |icon| icon_ids.must_include icon }
    end

    it 'includes these icons (reverse)' do
      @icons.each { |icon| Fixtures.icon_ids.must_include icon.id }
    end

    it 'does not includes these icons' do
      icons = %w(icon awesome)
      icons.each { |icon| @icons.wont_include icon }
    end
  end

  describe '#select!' do
    describe 'with `hdd`' do
      before do
        queries = %w(hdd)
        @icons = FontAwesome.new.select!(queries)
      end

      it { @icons.size.must_equal 1 }

      it 'must equal icon name' do
        icon_ids = @icons.map { |icon| icon.id }
        icon_ids.must_equal %w(hdd-o)
      end
    end

    describe 'with `left arr`' do
      before do
        queries = %w(left arr)
        @icons = FontAwesome.new.select!(queries)
        @icon_ids = %w(arrow-circle-left arrow-circle-o-left arrow-left long-arrow-left)
      end

      it { @icons.size.must_equal 4 }

      it 'must equal icon names' do
        icon_ids = @icons.map { |icon| icon.id }
        icon_ids.must_equal @icon_ids
      end
    end

    describe 'with `arr left` (reverse)' do
      before do
        queries = %w(arr left)
        @icons = FontAwesome.new.select!(queries)
        @icon_ids = %w(arrow-circle-left arrow-circle-o-left arrow-left long-arrow-left)
      end

      it { @icons.size.must_equal 4 }

      it 'must equal icon names' do
        icon_ids = @icons.map { |icon| icon.id }
        icon_ids.must_equal @icon_ids
      end
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
        @icon_ids = Fixtures.icon_ids
      end

      it { @icons.size.must_equal 409 }

      it 'must equal icon names' do
        icon_ids = @icons.map { |icon| icon.id }
        icon_ids.must_equal @icon_ids
      end
    end
  end

  describe '#item_hash' do
    before do
      @icon = FontAwesome::Icon.new('apple')
      @item_hash = FontAwesome.new.item_hash(@icon)
    end

    it { @item_hash[:uid].must_equal '' }
    it { @item_hash[:title].must_equal 'apple' }
    it { @item_hash[:subtitle].must_equal "Paste class name: fa-apple" }
    it { @item_hash[:arg].must_equal 'apple|||f179' }
    it { @item_hash[:icon][:type].must_equal 'default' }
    it { @item_hash[:icon][:name].must_equal "./icons/fa-apple.png" }
    it { @item_hash[:valid].must_equal 'yes' }
    it { @item_hash.size.must_equal 6 }
  end

  describe '#add_items' do
    before do
      feedback = Alfred::Core.new.feedback
      icon_ids = %w(beer cloud apple)
      icons = icon_ids.map { |name| FontAwesome::Icon.new(name) }
      @feedback = FontAwesome.new.add_items(feedback, icons)
    end

    it { @feedback.items.first.title.must_equal 'beer' }
    it { @feedback.items.last.title.must_equal 'apple' }
  end

  describe '#to_alfred' do
    before do
      alfred = Alfred::Core.new
      query = 'bookmark'
      xml = FontAwesome.new(query).to_alfred(alfred)
      @doc = REXML::Document.new(xml)
    end

    it { @doc.elements['items'].size.must_equal 2 }
    it { @doc.elements['items/item[1]/title'].text.must_equal 'bookmark' }
    it { @doc.elements['items/item[1]/arg'].text.must_equal 'bookmark|||f02e' }
    it { @doc.elements['items/item[1]/icon'].text.must_equal './icons/fa-bookmark.png' }
    it { @doc.elements['items/item[2]/title'].text.must_equal 'bookmark-o' }
    it { @doc.elements['items/item[2]/arg'].text.must_equal 'bookmark-o|||f097' }
    it { @doc.elements['items/item[2]/icon'].text.must_equal './icons/fa-bookmark-o.png' }
  end

  describe '::Icon' do
    describe '#initialize' do
      describe 'star-half-o (#detect_unicode_from_id)' do
        before { @icon = FontAwesome::Icon.new('star-half-o') }

        it { @icon.id.must_equal 'star-half-o' }
        it { @icon.unicode.must_equal 'f123' }
      end

      describe 'star-half-empty (#detect_unicode_from_aliases)' do
        before { @icon = FontAwesome::Icon.new('star-half-empty') }

        it { @icon.id.must_equal 'star-half-empty' }
        it { @icon.unicode.must_equal 'f123' }
      end

      it 'includes these icons' do
        Fixtures.icon_ids.each do |id|
          icon = FontAwesome::Icon.new(id)
          icon.id.must_equal id
          icon.unicode.wont_be_nil
        end
      end
    end
  end
end
