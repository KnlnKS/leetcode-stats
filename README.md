<p align="center">
 <h1 align="center">LeetCode Readme Stats</h1>
 <p align="center">Your LeetCode profile cards, now for your readmes!</p>
</p>

## Stats card

```md
[![LeetCode stats](https://leetcode-stats-six.vercel.app/KnlnKS)](https://github.com/KnlnKS/leetcode-stats)
```

[![LeetCode stats](https://leetcode-stats-six.vercel.app/KnlnKS)](https://github.com/KnlnKS/leetcode-stats)

## Submission graph

```md
[![LeetCode submissions](https://leetcode-stats-six.vercel.app/KnlnKS/graph)](https://github.com/KnlnKS/leetcode-stats)
```

[![LeetCode submissions](https://leetcode-stats-six.vercel.app/KnlnKS/graph)](https://github.com/KnlnKS/leetcode-stats)

## Options

| Parameter | Values | Default | |
| --- | --- | --- | --- |
| `theme` | `light`, `dark` | `light` | LeetCode's own two palettes |
| `font` | `default`, `system` | `default` | `default` embeds Inter so the card looks the same everywhere; `system` drops it and cuts the response from ~32 KB to ~3 KB |
| `cn` | flag | off | read from `leetcode.cn` instead |
| `width` | `300`–`2000` | `836` | graph card only |

```md
https://leetcode-stats-six.vercel.app/KnlnKS?theme=dark
https://leetcode-stats-six.vercel.app/KnlnKS/graph?theme=dark&width=500
https://leetcode-stats-six.vercel.app/KnlnKS?cn
```

## Matching GitHub's dark mode

`prefers-color-scheme` inside an SVG isn't honoured by GitHub, so pick the theme
from the markup instead:

```html
<picture>
  <source media="(prefers-color-scheme: dark)" srcset="https://leetcode-stats-six.vercel.app/KnlnKS?theme=dark">
  <img src="https://leetcode-stats-six.vercel.app/KnlnKS?theme=light">
</picture>
```

## Credits

Card design mirrors [LeetCode](https://leetcode.com). Typeface is
[Inter](https://rsms.me/inter/) by Rasmus Andersson, under the SIL Open Font
License 1.1 (`src/fonts/OFL.txt`).
