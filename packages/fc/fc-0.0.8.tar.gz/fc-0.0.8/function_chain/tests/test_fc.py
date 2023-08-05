from function_chain.fc import Fc

'''
Use `Pytest` as a test framework
'''


def test_map():
  l = [1, 2, 3, 4, 5, 6, 7]
  ml = Fc(l).map(lambda x: x + 1).done()
  assert ml == [2, 3, 4, 5, 6, 7, 8]


def test_filter():
  l = [1, 2, 3, 4, 5, 6, 7]
  ml = Fc(l).filter(lambda x: x > 3).filter(lambda x: x < 6).done()
  assert ml == [4, 5]


def test_sort():
  l = [7, 6, 5, 4, 3, 2, 1]
  ml = Fc(l).sort().done()
  assert ml == [1, 2, 3, 4, 5, 6, 7]


def test_resort():
  l = [1, 2, 3, 4, 5, 6, 7]
  ml = Fc(l).resort().done()
  assert ml == [7, 6, 5, 4, 3, 2, 1]


def test_len():
  l = [1, 2, 3, 4, 5, 6, 7]
  mll = Fc(l).filter(lambda x: x < 1).len()
  assert mll == 0

  mll = Fc(l).filter(lambda x: x > 4).len()
  assert mll == 3

  mll = Fc(l).filter(lambda x: x > 1 and x < 6).len()
  assert mll == 4


def test_reduce():
  l = [1, 2, 3, 4, 5, 6, 7]
  result = Fc(l).reduce(lambda x, y: x + y)
  assert result == 28


def test_iter():
  l = [1, 2, 3, 4, 5, 6, 7]
  nl = []
  for i in Fc(l).filter(lambda x: x > 2 and x < 5).iter():
    nl.append(i)
  assert nl == [3, 4]


def test_magic_iter():
  l = [1, 2, 3, 4, 5, 6, 7]
  nl = []
  for i in Fc(l).filter(lambda x: x > 2 and x < 5):
    nl.append(i)
  assert nl == [3, 4]