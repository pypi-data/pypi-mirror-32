"""Test base catlolzer functionalities."""

from catlolzer import (
    CatLolzer
)


def test_lolzing():
    catlolzer = CatLolzer(5)
    result = catlolzer.lolize("Miaow")
    assert result.shape[0] == 5
