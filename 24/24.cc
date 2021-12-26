#include <algorithm>
#include <array>
#include <cstdlib>
#include <fstream>
#include <functional>
#include <iostream>
#include <map>
#include <sstream>
#include <stdexcept>
#include <string>
#include <unordered_map>
#include <utility>
#include <vector>

namespace {

enum { INP = 0, ADD, MUL, DIV, MOD, EQL };
enum { W = 0, X, Y, Z };

auto reg(char r) {
  switch (r) {
  case 'w': return W;
  case 'x': return X;
  case 'y': return Y;
  case 'z': return Z;
  }
  std::exit(-1);
}

auto atoi(const std::string &s) {
  if (s[0] >= 'w' && s[0] <= 'z')
    return 100 + reg(s[0]);
  return std::stoi(s);
}

struct Op { int op, reg, value; };

auto opForLine(const std::string &line) {
  switch (line[1]) {
  case 'n': return Op{INP, reg(line[4]), 0};
  case 'd': return Op{ADD, reg(line[4]), atoi(line.substr(6))};
  case 'u': return Op{MUL, reg(line[4]), atoi(line.substr(6))};
  case 'i': return Op{DIV, reg(line[4]), atoi(line.substr(6))};
  case 'o': return Op{MOD, reg(line[4]), atoi(line.substr(6))};
  case 'q': return Op{EQL, reg(line[4]), atoi(line.substr(6))};
  }
  std::exit(-1);
}

auto readfile(const std::string &name) {
  auto rval = std::vector<std::vector<Op>>{};
  auto in = std::ifstream(name);
  auto line = std::string{};

  auto wip = std::vector<Op>{};
  while (std::getline(in, line)) {
    if (line[0] == '#')
      continue;
    auto op = opForLine(line);
    if (op.op == INP && !wip.empty()) {
      rval.push_back(wip);
      wip.clear();
    }
    wip.push_back(opForLine(line));
  }
  rval.push_back(wip);
  return rval;
}

auto run(const std::vector<Op> &input, int64_t value, int64_t z) {
  auto vars = std::array<int64_t, 4>{};
  vars[Z] = z;
  for (const auto &op : input) {
    switch (op.op) {
    case INP:
      vars[op.reg] = value;
      break;
    case ADD:
      vars[op.reg] += (op.value >= 100) ? vars[op.value - 100] : op.value;
      break;
    case MUL:
      vars[op.reg] *= (op.value >= 100) ? vars[op.value - 100] : op.value;
      break;
    case DIV:
      vars[op.reg] /= (op.value >= 100) ? vars[op.value - 100] : op.value;
      break;
    case MOD: {
      auto moduland = (op.value >= 100) ? vars[op.value - 100] : op.value;
      vars[op.reg] %= moduland;
      while (vars[op.reg] < 0)
        vars[op.reg] += moduland;
    } break;
    case EQL: {
      auto comp = (op.value >= 100) ? vars[op.value - 100] : op.value;
      vars[op.reg] = (vars[op.reg] == comp) ? 1 : 0;
    } break;
    }
  }
  return vars[Z];
}
//
template <class F>
auto evaluate(std::vector<std::vector<Op>> &ops, F resolver) {
  auto results = std::unordered_map<int64_t, int64_t>{};
  results[0] = 0;

  auto working = std::unordered_map<int64_t, int64_t>{};
  for (auto level = 0; level <= 13; level++) {
    for (const auto &pair : results) {
      for (auto x = 1; x <= 9; x++) {
        if (pair.first > 1000000)
          continue;
        auto e = run(ops[level], x, pair.first);
        auto val = pair.second * 10 + x;
        if (auto f = working.find(e); f != working.end()) {
          f->second = resolver(f->second, val);
        } else {
          working[e] = val;
        }
      }
    }
    std::swap(working, results);
    working.clear();
  }
  std::cout << results.at(0) << '\n';
}

} // namespace

int main() {
  auto ops = readfile("input");

  evaluate(ops, [](int64_t a, int64_t b) { return std::max(a, b); }); // 99598963999971
  evaluate(ops, [](int64_t a, int64_t b) { return std::min(a, b); }); // 93151411711211
}
