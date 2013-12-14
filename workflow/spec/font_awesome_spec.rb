# encoding: utf-8

require File.expand_path('spec_helper', File.dirname(__FILE__))

describe FontAwesome do
  it 'does not cause an error' do
    actual = require('bundle/bundler/setup')
    expect(actual).to be false
  end

  describe '.to_character_reference' do
    it 'returns the character reference' do
      actual = FontAwesome.to_character_reference('f000')
      expect(actual).to eq('')

      actual = FontAwesome.to_character_reference('f17b')
      expect(actual).to eq('')
    end

    it 'does not returns the character reference' do
      actual = FontAwesome.to_character_reference('f001')
      expect(actual).not_to eq('')
    end
  end

  describe '#icons' do
    let(:icons) { FontAwesome.new.icons }

    it 'returns 409' do
      expect(icons.size).to eq(409)
    end

    it 'returns "adjust"' do
      expect(icons.first.id).to eq('adjust')
    end

    it 'returns "youtube-square"' do
      expect(icons.last.id).to eq('youtube-square')
    end

    it 'includes these icons' do
      icon_ids = icons.map { |icon| icon.id }
      Fixtures.icon_ids.each { |icon| expect(icon_ids).to be_include(icon) }
    end

    it 'includes these icons (reverse)' do
      icons.each { |icon| expect(Fixtures.icon_ids).to be_include(icon.id) }
    end

    it 'does not includes these icons' do
      expectation = %w(icon awesome)
      expectation.each { |icon| expect(icons).not_to be_include(icon) }
    end
  end

  describe '#select!' do
    context 'with "hdd"' do
      let(:icons) { FontAwesome.new.select!(%w(hdd)) }

      it 'returns 1' do
        expect(icons.size).to eq(1)
      end

      it 'must equal icon name' do
        icon_ids = icons.map { |icon| icon.id }
        expect(icon_ids).to eq(%w(hdd-o))
      end
    end

    context 'with "left arr"' do
      let(:icons) { FontAwesome.new.select!(%w(left arr)) }

      it 'returns 4' do
        expect(icons.size).to eq(4)
      end

      it 'must equal icon names' do
        icon_ids = icons.map { |icon| icon.id }
        expectation = %w(
          arrow-circle-left
          arrow-circle-o-left
          arrow-left long-arrow-left
        )
        expect(icon_ids).to eq(expectation)
      end
    end

    context 'with "arr left" (reverse)' do
      let(:icons) { FontAwesome.new.select!(%w(arr left)) }

      it 'returns 4' do
        expect(icons.size).to eq(4)
      end

      it 'must equal icon names' do
        icon_ids = icons.map { |icon| icon.id }
        expectation = %w(
          arrow-circle-left
          arrow-circle-o-left
          arrow-left long-arrow-left
        )
        expect(icon_ids).to eq(expectation)
      end
    end

    context 'with "icon" (does not match)' do
      let(:icons) { FontAwesome.new.select!(%w(icon)) }

      it 'returns a empty array' do
        expect(icons).to eq([])
      end
    end

    context 'with unknown arguments' do
      let(:icons) { FontAwesome.new.select!([]) }

      it 'returns 409' do
        expect(icons.size).to eq(409)
      end

      it 'must equal icon names' do
        icon_ids = icons.map { |icon| icon.id }
        expect(icon_ids).to eq(Fixtures.icon_ids)
      end
    end
  end

  describe '#item_hash' do
    let(:item_hash) do
      icon = FontAwesome::Icon.new('apple')
      FontAwesome.new.item_hash(icon)
    end

    it 'returns 6' do
      expect(item_hash.size).to eq(6)
    end

    it 'must equal hash values' do
      expect(item_hash[:uid]).to eq('')
      expect(item_hash[:title]).to eq('apple')
      expect(item_hash[:subtitle]).to eq('Paste class name: fa-apple')
      expect(item_hash[:arg]).to eq('apple|||f179')
      expect(item_hash[:icon][:type]).to eq('default')
      expect(item_hash[:icon][:name]).to eq('./icons/fa-apple.png')
      expect(item_hash[:valid]).to eq('yes')
    end
  end

  describe '#item_xml' do
    let(:item_xml) do
      icon = FontAwesome::Icon.new('apple')
      item_hash = FontAwesome.new.item_hash(icon)
      FontAwesome.new.item_xml(item_hash)
    end

    it 'returns the XML' do
      expectation = <<-XML
<item arg="apple|||f179" uid="">
  <title>apple</title>
  <subtitle>Paste class name: fa-apple</subtitle>
  <icon>./icons/fa-apple.png</icon>
</item>
      XML
      expect(item_xml).to eq(expectation)
    end
  end

  describe '#to_alfred' do
    let(:doc) do
      queries = ['bookmark']
      xml = FontAwesome.new(queries).to_alfred
      REXML::Document.new(xml)
      # TODO: mute puts
    end

    it 'returns 2' do
      expect(doc.elements['items'].elements.size).to eq(2)
    end

    it 'must equal XML elements' do
      expect(doc.elements['items/item[1]'].attributes['arg']).to \
        eq('bookmark|||f02e')
      expect(doc.elements['items/item[1]/title'].text).to eq('bookmark')
      expect(doc.elements['items/item[1]/icon'].text).to \
        eq('./icons/fa-bookmark.png')
      expect(doc.elements['items/item[2]'].attributes['arg']).to \
        eq('bookmark-o|||f097')
      expect(doc.elements['items/item[2]/title'].text).to eq('bookmark-o')
      expect(doc.elements['items/item[2]/icon'].text).to \
        eq('./icons/fa-bookmark-o.png')
    end

    it 'must equal $stdout (test for puts)' do
      expectation = <<-XML
<?xml version='1.0'?>
<items>
<item arg="bookmark|||f02e" uid="">
  <title>bookmark</title>
  <subtitle>Paste class name: fa-bookmark</subtitle>
  <icon>./icons/fa-bookmark.png</icon>
</item>
<item arg="bookmark-o|||f097" uid="">
  <title>bookmark-o</title>
  <subtitle>Paste class name: fa-bookmark-o</subtitle>
  <icon>./icons/fa-bookmark-o.png</icon>
</item>
</items>
      XML

      actual = capture(:stdout) { FontAwesome.new(['bookmark']).to_alfred }
      expect(actual).to eq(expectation)
    end
  end

  describe '::Icon' do
    describe '#initialize' do
      context 'star-half-o (#detect_unicode_from_id)' do
        let(:icon) { FontAwesome::Icon.new('star-half-o') }

        it 'returns "star-half-o"' do
          expect(icon.id).to eq('star-half-o')
        end

        it 'returns "f123"' do
          expect(icon.unicode).to eq('f123')
        end
      end

      context 'star-half-empty (#detect_unicode_from_aliases)' do
        let(:icon) { FontAwesome::Icon.new('star-half-empty') }

        it 'returns "star-half-o"' do
          expect(icon.id).to eq('star-half-empty')
        end

        it 'returns "f123"' do
          expect(icon.unicode).to eq('f123')
        end
      end

      it 'includes these icons' do
        Fixtures.icon_ids.each do |id|
          icon = FontAwesome::Icon.new(id)
          expect(icon.id).to eq(id)
          expect(icon.unicode).not_to be_nil
        end
      end
    end
  end
end
