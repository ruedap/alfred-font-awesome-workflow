# encoding: utf-8

require File.expand_path('spec_helper', File.dirname(__FILE__))

describe FontAwesome do
  let(:config_file_path) { described_class::CONFIG_FILE_PATH }
  let(:glob_icons) { described_class.glob_icons }
  let(:version) { described_class::VERSION }

  it 'does not cause an error' do
    actual = require('bundle/bundler/setup')
    expect(actual).to be false
  end

  describe '.argv' do
    it 'returns the OpenStruct object for ARGV' do
      argv = ['adjust|||f042']
      actual = described_class.argv(argv)
      expect(actual.icon_id).to eq('adjust')
      expect(actual.icon_unicode).to eq('f042')
    end
  end

  describe '.css_class_name' do
    it 'returns the CSS class name' do
      actual = described_class.css_class_name('adjust')
      expect(actual).to eq('fa-adjust')
    end
  end

  describe '.character_reference' do
    it 'returns the character reference' do
      actual = described_class.character_reference('f000')
      expect(actual).to eq('')

      actual = described_class.character_reference('f17b')
      expect(actual).to eq('')
    end

    it 'does not returns the character reference' do
      actual = described_class.character_reference('f001')
      expect(actual).not_to eq('')
    end
  end

  describe '.glob_icons' do
    it 'returns 593' do
      expect(glob_icons.size).to eq(593)
    end

    it 'returns "adjust"' do
      allow(described_class).to receive(:load_config).and_return({})
      expect(glob_icons.first.id).to eq('adjust')
    end

    it 'returns "youtube-square"' do
      expect(glob_icons.last.id).to eq('youtube-square')
    end

    it 'includes these icons' do
      icon_ids = glob_icons.map { |icon| icon.id }
      Fixtures.icon_ids.each { |icon| expect(icon_ids).to be_include(icon) }
    end

    it 'includes these icons (reverse)' do
      glob_icons.each { |icon| expect(Fixtures.icon_ids).to be_include(icon.id) }
    end

    it 'does not include these icons' do
      expectation = %w(icon awesome)
      expectation.each { |icon| expect(glob_icons).not_to be_include(icon) }
    end
  end

  describe '.load_config' do
    context 'when does not exist config.yml' do
      it 'returns config.yaml and contains version number' do
        FileUtils.rm(config_file_path) if File.exist?(config_file_path)

        expect(File.exist?(config_file_path)).to be_falsy
        expect(described_class.load_config['version']).to eq(version)
        expect(File.exist?(config_file_path)).to be_truthy
      end
    end

    context 'when exists config.yml' do
      it 'returns config.yaml and contains version number' do
        unless File.exist?(config_file_path)
          open(config_file_path, 'w') { |f| YAML.dump({ 'version' => version }, f) }
        end

        expect(File.exist?(config_file_path)).to be_truthy
        expect(described_class.load_config['version']).to eq(version)
        expect(File.exist?(config_file_path)).to be_truthy
      end
    end

    context 'when raises an error' do
      before do
        allow(YAML).to receive(:load_file).and_raise
      end

      it 'returns nil' do
        expect(described_class.load_config).to be_nil
      end

      it 'exists config.yml' do
        expect(File.exist?(config_file_path)).to be_truthy
      end
    end

    context 'when bump version number' do
      it 'contains old version number' do
        FileUtils.rm(config_file_path) if File.exist?(config_file_path)

        stub_const("#{described_class}::VERSION", '1.0.0.0')
        expect(described_class.load_config['version']).to eq('1.0.0.0')

        stub_const("#{described_class}::VERSION", '2.0.0.0')
        expect(described_class.load_config['version']).to eq('1.0.0.0')
      end
    end
  end

  describe '.save_config_of_recent_icons' do
    context 'when does not exist config.yml' do
      before do
        FileUtils.rm(config_file_path) if File.exist?(config_file_path)
      end

      it 'exists config.yml' do
        expect(File.exist?(config_file_path)).to be_falsy
        actual = described_class.save_config_of_recent_icons('apple')
        expect(actual.instance_of?(File)).to be_truthy
        expect(File.exist?(config_file_path)).to be_truthy
      end

      it 'contains recent icons in config.yml' do
        described_class.save_config_of_recent_icons('twitter')
        described_class.save_config_of_recent_icons('github')
        described_class.save_config_of_recent_icons('apple')

        actual = described_class.load_config['recent_icons']
        expect(actual).to eq(%w(apple github twitter))
      end
    end

    context 'when exists config.yml' do
      before do
        unless File.exist?(config_file_path)
          open(config_file_path, 'w') { |f| YAML.dump({ 'version' => version }, f) }
        end
      end

      it 'exists config.yml' do
        expect(File.exist?(config_file_path)).to be_truthy
        actual = described_class.save_config_of_recent_icons('apple')
        expect(actual.instance_of?(File)).to be_truthy
        expect(File.exist?(config_file_path)).to be_truthy
      end

      it 'contains recent icons in config.yml' do
        described_class.save_config_of_recent_icons('twitter')
        described_class.save_config_of_recent_icons('github')
        described_class.save_config_of_recent_icons('apple')

        actual = described_class.load_config['recent_icons']
        expect(actual).to eq(%w(apple github twitter))
      end
    end

    context 'when raises an error' do
      it 'returns nil' do
        allow(described_class).to receive(:dump_config).and_raise

        actual = described_class.save_config_of_recent_icons('apple')
        expect(actual).to be_nil
      end
    end

    context 'when bump version number' do
      it 'contains new version number' do
        FileUtils.rm(config_file_path) if File.exist?(config_file_path)

        stub_const("#{described_class}::VERSION", '1.0.0.0')
        described_class.save_config_of_recent_icons('apple')
        expect(described_class.load_config['version']).to eq('1.0.0.0')

        stub_const("#{described_class}::VERSION", '2.0.0.0')
        described_class.save_config_of_recent_icons('apple')
        expect(described_class.load_config['version']).to eq('2.0.0.0')
      end
    end
  end

  describe '.url' do
    it 'returns the Font Awesome URL' do
      actual = described_class.url('adjust')
      expect(actual).to eq('http://fontawesome.io/icon/adjust/')
    end
  end

  describe '#initialize' do
    it 'assigns @selected_icons' do
      fa = described_class.new
      actual = fa.instance_variable_get(:@selected_icons)
      expect(actual).not_to be_nil
    end
  end

  describe '#select!' do
    context 'with "hdd"' do
      let(:icons) { described_class.new.select!(%w(hdd), glob_icons) }

      it 'returns 1' do
        expect(icons.size).to eq(1)
      end

      it 'must equal icon name' do
        icon_ids = icons.map { |icon| icon.id }
        expect(icon_ids).to eq(%w(hdd-o))
      end
    end

    context 'with "left arr"' do
      let(:icons) { described_class.new.select!(%w(left arr), glob_icons) }

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
      let(:icons) { described_class.new.select!(%w(arr left), glob_icons) }

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

    context 'with "icons" (does not match)' do
      let(:icons) { described_class.new.select!(%w(icons), glob_icons ) }

      it 'returns an empty array' do
        expect(icons).to eq([])
      end
    end

    context 'with unknown arguments' do
      let(:icons) { described_class.new.select!([], glob_icons) }

      it 'returns 593' do
        expect(icons.size).to eq(593)
      end

      it 'must equal icon names' do
        allow(described_class).to receive(:load_config).and_return({})
        icon_ids = icons.map { |icon| icon.id }
        expect(icon_ids).to eq(Fixtures.icon_ids)
      end
    end

    context 'with "taxi"' do  # for ver.4.1.0
      let(:icons) { described_class.new.select!(%w(taxi), glob_icons) }

      it 'must equal icon name' do
        icon_ids = icons.map { |icon| icon.id }
        expect(icon_ids).to eq(%w(taxi))
      end
    end

    context 'with "angellist"' do  # for ver.4.2.0
      let(:icons) { described_class.new.select!(%w(angellist), glob_icons) }

      it 'must equal icon name' do
        icon_ids = icons.map { |icon| icon.id }
        expect(icon_ids).to eq(%w(angellist))
      end
    end

    context 'with "medium"' do  # for ver.4.3.0
      let(:icons) { described_class.new.select!(%w(medium), glob_icons) }

      it 'must equal icon name' do
        icon_ids = icons.map { |icon| icon.id }
        expect(icon_ids).to eq(%w(medium))
      end
    end
  end

  describe '#sort_by_recent_icons' do
    context 'when does not exist recent_icons property in config.yml' do
      before do
        config_yaml = {
          'version' => '0.0.0.0'
        }
        allow(described_class).to receive(:load_config).and_return(config_yaml)
      end

      it 'returns alphabetically icons' do
        actual = described_class.new.sort_by_recent_icons(glob_icons)
        expect(actual.first.id).to eq('adjust')
        expect(actual[100].id).to eq('certificate')
        expect(actual[200].id).to eq('file-code-o')
        expect(actual.last.id).to eq('youtube-square')
      end
    end

    context 'when exists recent_icons property in config.yml' do
      before do
        config_yaml = {
          'version' => '0.0.0.0',
          'recent_icons' => %w(apple github twitter)
        }
        allow(described_class).to receive(:load_config).and_return(config_yaml)
      end

      it 'returns icons of recently used order' do
        actual = described_class.new.sort_by_recent_icons(glob_icons)
        expect(actual.first.id).to eq('apple')
        expect(actual[1].id).to eq('github')
        expect(actual[2].id).to eq('twitter')
        expect(actual[100].id).to eq('cc-stripe')
        expect(actual[200].id).to eq('file-archive-o')
        expect(actual.last.id).to eq('youtube-square')
      end
    end
  end

  describe '#item_hash' do
    let(:item_hash) do
      icon = described_class::Icon.new('apple')
      described_class.new.item_hash(icon)
    end

    it 'returns 6' do
      expect(item_hash.size).to eq(6)
    end

    it 'must equal hash values' do
      expect(item_hash[:uid]).to eq('apple')
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
      icon = described_class::Icon.new('apple')
      item_hash = described_class.new.item_hash(icon)
      described_class.new.item_xml(item_hash)
    end

    it 'returns the XML' do
      Timecop.freeze(Time.now) do
        expectation = <<-XML
<item arg="apple|||f179" uid="#{Time.now.to_i}-apple">
<title>apple</title>
<subtitle>Paste class name: fa-apple</subtitle>
<icon>./icons/fa-apple.png</icon>
</item>
        XML
        expect(item_xml).to eq(expectation)
      end
    end
  end

  describe '#to_alfred' do
    let(:doc) do
      queries = ['bookmark']
      fa = described_class.new(queries)
      allow(fa).to receive(:puts) # Mute puts
      xml = fa.to_alfred
      REXML::Document.new(xml)
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
      Timecop.freeze(Time.now) do
        expectation = <<-XML
<?xml version='1.0'?><items><item arg="bookmark|||f02e" uid="#{Time.now.to_i}-bookmark"><title>bookmark</title><subtitle>Paste class name: fa-bookmark</subtitle><icon>./icons/fa-bookmark.png</icon></item><item arg="bookmark-o|||f097" uid="#{Time.now.to_i}-bookmark-o"><title>bookmark-o</title><subtitle>Paste class name: fa-bookmark-o</subtitle><icon>./icons/fa-bookmark-o.png</icon></item></items>
        XML

        actual = capture(:stdout) { described_class.new(['bookmark']).to_alfred }
        expect(actual).to eq(expectation)
      end
    end
  end

  describe '::Icon' do
    describe '#initialize' do
      context 'star-half-o (#detect_unicode_from_id)' do
        let(:icon) { described_class::Icon.new('star-half-o') }

        it 'returns "star-half-o"' do
          expect(icon.id).to eq('star-half-o')
        end

        it 'returns "f123"' do
          expect(icon.unicode).to eq('f123')
        end
      end

      context 'star-half-empty (#detect_unicode_from_aliases)' do
        let(:icon) { described_class::Icon.new('star-half-empty') }

        it 'returns "star-half-o"' do
          expect(icon.id).to eq('star-half-empty')
        end

        it 'returns "f123"' do
          expect(icon.unicode).to eq('f123')
        end
      end

      it 'includes these icons' do
        Fixtures.icon_ids.each do |id|
          icon = described_class::Icon.new(id)
          expect(icon.id).to eq(id)
          expect(icon.unicode).not_to be_nil
        end
      end
    end
  end
end
