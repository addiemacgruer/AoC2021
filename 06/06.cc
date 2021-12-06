#include <array>
#include <fstream>
#include <iostream>

namespace {

using FishMap = std::array<uint64_t, 9>;

auto read_file(const std::string &name) {
  auto ih = std::ifstream{name};
  auto line = std::string{};
  ih >> line;
  auto rval = FishMap();
  for (auto c : line) {
    if (c == ',')
      continue;
    rval[c - '0']++;
  }
  return rval;
}

auto fishcount(const FishMap &fish) {
  auto rval = uint64_t{};
  for (const auto &e : fish)
    rval += e;
  return rval;
}

auto next_day(const FishMap &fish) {
  auto rval = FishMap();
  rval[6] = fish[0];
  rval[8] = fish[0];
  for (int i = 1; i <= 8; ++i)
    rval[i - 1] += fish[i];
  return rval;
}

} // namespace

int main() {
  auto fish = read_file("input");
  for (auto i = 0; i < 256; ++i) {
    fish = next_day(fish);
    if (i == 80)
      std::cout << fishcount(fish) << '\n'; // 380243
  }
  std::cout << fishcount(fish) << '\n'; // 1708791884591
}
