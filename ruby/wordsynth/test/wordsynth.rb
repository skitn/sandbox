require 'minitest/autorun'
require_relative '../lib/effects'
require_relative '../lib/wordsynth'

class WordSynthTest < Minitest::Test
  def test_play
    assert WordSynth
    assert Effects
  end
end
