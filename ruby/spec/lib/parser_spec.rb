require "spec_helper"
require "parser"

describe "main" do
    context "source" do
        it "downloads a page from Angry metal guy" do
            allow(URI).to receive(:open).and_return(File.new("spec/fixtures/page.html"))
            expect(::Nokogiri::HTML::Document).to receive(:parse).and_call_original
            main
        end
    end

    context "format" do
        before do
            allow(URI).to receive(:open).and_return(File.new("spec/fixtures/page.html"))
        end
        it "returns the right band" do
            res = main
            expect(["//Wilderun – Epigone ", "//Killing Joke – Pylon ", "//Wilderun – Veil of Imagination ", "//Fellowship – The Saberlight Chronicles ", "//Thränenkind – King Apathy ", "//Pyrrhon – What Passes for Survival ", "//Impure Wilhelmina – Antidote ", "//Ad Nauseam – Imperative Imperceptible Impulse "]).to include(res)
        end
    end

end