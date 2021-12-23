#include <iostream>
#include <map>
#include <stdexcept>
#include <vector>

namespace {

struct Point {
  int x, y;

  auto operator<(const Point &o) const {
    if (x == o.x)
      return y < o.y;
    return x < o.x;
  }

  auto operator==(const Point &o) const { return x == o.x && y == o.y; }
  auto operator!=(const Point &o) const { return x != o.x || y != o.y; }
};

auto costs = std::map<char, int>{{'A', 1}, {'B', 10}, {'C', 100}, {'D', 1000}};
auto slots = std::map<char, int>{{'A', 2}, {'B', 4}, {'C', 6}, {'D', 8}};
auto targets = std::vector<int>{0, 1, 3, 5, 7, 9, 10};

auto cost(const Point &a, const Point &b, char type) {
  auto dx = std::abs(a.x - b.x);
  auto dy = (a.y != 0 && b.y != 0) ? a.y + b.y : std::abs(a.y - b.y);
  return costs[type] * (dx + dy);
}

using Positions = std::map<Point, char>;

struct Hallway {
  Positions points{};
  int total = 0;

  auto height() const { return points.size() == 16 ? 5 : 3; }

  auto print() const {
    for (auto y = int{}; y < height(); ++y) {
      for (auto x = 0; x < 11; ++x) {
        auto p = points.find(Point{x, y});
        std::cout << (p != points.end() ? p->second : '.');
      }
      std::cout << '\n';
    }
  }

  auto is_win() const {
    for (const auto &e : points)
      if (e.first.x != slots.at(e.second))
        return false;
    return true;
  }

  auto move(const Point &a, const Point &b) {
    auto t = points.at(a);
    auto c = cost(a, b, t);
    points[b] = t;
    points.erase(a);
    total += c;
  }

  auto at(int x, int y) const {
    auto a = points.find(Point{x, y});
    return (a == points.end()) ? ' ' : a->second;
  }

  auto can_move_down(int start) const {
    auto type = at(start, 0);
    auto target = slots.at(type);
    auto smallest = height();
    for (const auto &e : points) {
      if (e.first.x == target && e.second != type)
        return -1;
      if (e.first.x == target && e.second == type)
        smallest = std::min(e.first.y, smallest);
    }
    for (auto x = std::min(start, target); x <= std::max(start, target); ++x) {
      if (x == start)
        continue;
      if (at(x, 0) != ' ')
        return -1;
    }
    return smallest - 1;
  }

  auto can_move_up(int start, int end) const {
    for (auto x = std::min(start, end); x <= std::max(start, end); ++x) {
      if (at(x, 0) != ' ')
        return -1;
    }
    auto smallest = height() + 1;
    auto finished = true;
    for (const auto &e : points) {
      if (e.first.x == start && e.first.y < smallest)
        smallest = std::min(smallest, e.first.y);
      if (slots.at(e.second) != start)
        finished = false;
    }
    if (finished || smallest == height() + 1)
      return -1;
    return smallest;
  }

  auto possible_moves() const {
    auto rval = std::vector<Hallway>();
    // top row
    for (auto x = 0; x < 11; ++x) {
      if (auto a = at(x, 0); a != ' ') {
        if (auto y = can_move_down(x); y > 0) {
          auto copy = *this;
          copy.move(Point{x, 0}, Point{slots.at(a), y});
          rval.push_back(std::move(copy));
        }
      }
    }
    // bottom rows
    for (auto x = 2; x <= 8; x += 2) {
      for (auto t : targets) {
        if (auto y = can_move_up(x, t); y > 0) {
          auto copy = *this;
          copy.move(Point{x, y}, Point{t, 0});
          rval.push_back(std::move(copy));
        }
      }
    }
    return rval;
  }
};

auto hallway(const std::string &v) {
  auto rval = Hallway{};
  for (auto i = int{}; i < v.size(); ++i)
    rval.points[Point{(i % 4 + 1) * 2, i / 4 + 1}] = v[i];
  return rval;
}

auto puzzle = hallway("DAADCCBB");
auto ptwo = hallway("DAADDCBADBACCCBB");

auto solve(const Hallway &h) {
  auto current = std::vector<Hallway>();
  current.push_back(h);
  while (true) {
    auto possibles = std::map<Positions, int>();
    for (const auto &c : current) {
      auto poss = c.possible_moves();
      for (const auto &p : poss) {
        auto i = possibles.find(p.points);
        if (i == possibles.end()) {
          possibles[p.points] = p.total;
        } else {
          i->second = std::min(i->second, p.total);
        }
      }
    }
    current.clear();
    for (auto &p : possibles) {
      current.push_back(Hallway{p.first, p.second});
    }
    auto done = false;
    for (auto &c : current) {
      if (c.is_win()) {
        done = true;
        std::cout << c.total << '\n';
      }
    }
    if (done)
      break;
  }
}

} // namespace

int main() {
  solve(puzzle);
  solve(ptwo);
}
