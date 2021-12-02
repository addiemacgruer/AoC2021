#include <fstream>
#include <iostream>
#include <vector>

template <class T> auto read_file(const std::string &name) {
  auto file = std::ifstream(name);
  auto items = std::vector<T>{};
  auto item = T{};
  while (file >> item)
    items.emplace_back(item);
  return items;
}

template <class T> auto interval(const std::vector<T> &gaps, T gap) {
  auto greater = 0;
  for (auto i = size_t{}, end = gaps.size() - gap; i < end; ++i)
    if (gaps[i + gap] > gaps[i])
      ++greater;
  return greater;
}

int main() {
  auto items = read_file<int>("input");
  std::cout << interval(items, 1) << '\n'  // 1548
            << interval(items, 3) << '\n'; // 1589
}
