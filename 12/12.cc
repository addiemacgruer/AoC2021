#include <algorithm>
#include <fstream>
#include <iostream>
#include <map>
#include <vector>

namespace {

constexpr auto START = 0;
constexpr auto END = 1000;
struct Path { std::string start, end; };
using CaveList = std::vector<int>;
using CaveMap = std::map<int, CaveList>;
using PathList = std::vector<int>;

auto &operator>>(std::istream &in, Path &path) { // "createPath"
  auto line = std::string{};
  in >> line;
  path.start = line.substr(0, line.find('-'));
  path.end = line.substr(line.find('-') + 1);
  return in;
}

auto is_small(const std::string &cave) {
  return cave[0] >= 'a' && cave[0] <= 'z';
}

auto weight(const std::string &cave) { return is_small(cave) ? -1 : 1; }

auto read_file(const std::string &name) {
  auto file = std::ifstream{name};

  auto caveMap = CaveMap{};
  auto names = std::map<std::string, int>{};
  names["start"] = START;
  names["end"] = END;
  auto next = 1;
  auto path = Path{};
  while (file >> path) {
    if (names.find(path.start) == names.end()) {
      names[path.start] = weight(path.start) * next++;
    }
    if (names.find(path.end) == names.end()) {
      names[path.end] = weight(path.end) * next++;
    }
    caveMap[names[path.start]].push_back(names[path.end]);
    caveMap[names[path.end]].push_back(names[path.start]);
  }
  return caveMap;
}

auto partOne(int next, const PathList &work) {
  return next <= 0 && std::find(work.begin(), work.end(), next) != work.end();
}

auto partTwo(int next, const PathList &work) {
  if (next == 0) {
    return true;
  }
  if (!partOne(next, work)) {
    return false;
  }
  for (auto i = 1; i < work.size(); ++i) {
    if (work[i] > 0) {
      continue;
    }
    for (auto j = i + 1; j < work.size(); ++j) {
      if (work[i] == work[j]) {
        return true;
      }
    }
  }
  return false;
}

template <class Rejecter>
auto pathCount(const CaveMap &caveMap, Rejecter rejecter) {
  auto workingPaths = std::vector<PathList>{};
  auto rval = 0;
  workingPaths.push_back(PathList{START});
  while (workingPaths.size()) {
    auto work = workingPaths.back();
    workingPaths.pop_back();
    for (auto &next : caveMap.at(work.back())) {
      if (rejecter(next, work)) {
        continue;
      }
      if (next == END) {
        ++rval;
      } else {
        auto prospective = work;
        prospective.push_back(next);
        workingPaths.push_back(prospective);
      }
    }
  }
  return rval;
}

} // namespace

auto main() -> int {
  auto caveMap = read_file("input.test");
  std::cout << pathCount(caveMap, partOne) << '\n'  // 10
            << pathCount(caveMap, partTwo) << '\n'; // 36
  caveMap = read_file("input.test2");
  std::cout << pathCount(caveMap, partOne) << '\n'  // 19
            << pathCount(caveMap, partTwo) << '\n'; // 103
  caveMap = read_file("input.test3");
  std::cout << pathCount(caveMap, partOne) << '\n'  // 226
            << pathCount(caveMap, partTwo) << '\n'; // 3509
  caveMap = read_file("input");
  std::cout << pathCount(caveMap, partOne) << '\n'  // 4304
            << pathCount(caveMap, partTwo) << '\n'; // 118242
}
