require "nokogiri"
require "open-uri"
require "debug"
require "ostruct"

def main
    to_mpv(album.title)
end

def url
    "https://www.angrymetalguy.com/tag/50/"
end

private def doc
    @doc ||= Nokogiri::HTML(URI.open(url))
end

private def all_reviews
    doc.xpath("//article[contains(@class, 'tag-review')]")
end

private def albums
   @albums ||= begin
        all_reviews.map { |album| 
            title = album.xpath(".//h2[contains(@class, 'entry-title')]/a").text
            title.gsub!(/review/i, '')
            OpenStruct.new(:title => title)
        }
    end
end

private def album
    albums.sample(1).first
end

private def to_mpv(album)
    "//#{album}"
end

