#include <fstream>
#include <functional>
#include <iostream>
#include <limits>
#include <map>
#include <tuple>
#include <unordered_map>
#include <vector>

namespace {

using Subs = std::map<std::string, std::string>;
struct Step { std::string pair; int level; };
bool operator==(const Step &a, const Step &b) {
  return a.level == b.level && a.pair == b.pair;
}
using Dist = std::map<std::string::value_type, uint64_t>;

} // namespace

namespace std {
template <> struct hash<Step> {
  size_t operator()(const Step &x) const {
    return std::hash<decltype(x.pair)>()(x.pair) ^
           std::hash<decltype(x.level)>()(x.level);
  }
};
} // namespace std

namespace {

auto read_file(const std::string &name) {
  auto file = std::ifstream{name};
  auto molecule = std::string{};
  file >> molecule;
  auto line = std::string{};
  auto subs = Subs{};
  while (std::getline(file, line)) {
    if (line.empty()) continue;
    subs[line.substr(0, line.find(' '))] = line.substr(line.rfind(' ') + 1);
  }
  return std::make_tuple(molecule, subs);
}

auto operator+(Dist copy, const Dist &b) {
  for (auto &kv : b) copy[kv.first] += kv.second;
  return copy;
}

auto cache = std::unordered_map<Step, Dist>{};

auto evaluate_pair(const std::string &pair, const Subs &subs, int limit) {
  auto next = Step{pair, limit};
  if (auto existing = cache.find(next); existing != cache.end()) {
    return existing->second;
  }
  if (limit == 0) {
    auto rval = Dist{std::make_pair(pair[0], 1)};
    cache[next] = rval;
    return rval;
  }
  auto s = std::string{subs.find(pair)->second};
  auto mapA = evaluate_pair(pair.substr(0, 1) + s, subs, limit - 1);
  auto mapB = evaluate_pair(s + pair.substr(1, 1), subs, limit - 1);
  auto rval = mapA + mapB;
  cache[next] = rval;
  return rval;
}

auto min_max(const Dist &dist) {
  auto min = std::numeric_limits<Dist::mapped_type>::max();
  auto max = Dist::mapped_type{};
  for (auto &kv : dist) {
    min = std::min(min, kv.second);
    max = std::max(max, kv.second);
  }
  return std::make_tuple(min, max);
}

auto evaluate_molecule(const std::string &mol, const Subs &subs, int limit) {
  auto result = Dist{};
  for (auto i = size_t{}; i < mol.size() - 1; ++i) {
    result = result + evaluate_pair(mol.substr(i, 2), subs, limit);
  }
  result[mol[mol.size() - 1]]++;
  auto [min, max] = min_max(result);
  return max - min;
}

} // namespace

int main() {
  auto [mol, subs] = read_file("input");
  std::cout << evaluate_molecule(mol, subs, 10) << '\n'  // 2768
            << evaluate_molecule(mol, subs, 40) << '\n'; // 2914365137499
}
