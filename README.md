<p align="center">
 <h1 align="center">LeetCode Readme Stats</h1>
 <p align="center">Dynamically generated LeetCode stats for your readmes!</p>
</p>

## How To Use
Copy-paste this into your markdown content, and that's it. Simple!
Change the `?username=` value to your LeetCode username.

```md
<a href="https://github.com/KnlnKS/leetcode-stats">
  <img alt="LeetCode Stat Card" src="https://apu5rh8gxk.execute-api.us-east-1.amazonaws.com/default/leetcode-stats?username=KnlnKS" width="400"/>
</a>
```

Which will appear as:

<a href="https://github.com/KnlnKS/leetcode-stats">
  <img alt="LeetCode Stat Card" src="https://apu5rh8gxk.execute-api.us-east-1.amazonaws.com/default/leetcode-stats?username=KnlnKS" width="400"/>
</a>

## Themes
You can customize your card with themes! There are 2 themes available.
light (default) and dark. To use them, simply add `&theme=` to the
end of the url with the theme you'd like to use.
Ex.

```md
<a href="https://github.com/KnlnKS/leetcode-stats">
  <img alt="LeetCode Stat Card" src="https://apu5rh8gxk.execute-api.us-east-1.amazonaws.com/default/leetcode-stats?username=KnlnKS" width="400"/>
</a>
```

Which will appear as:

<a href="https://github.com/KnlnKS/leetcode-stats">
  <img alt="LeetCode Stat Card" src="https://apu5rh8gxk.execute-api.us-east-1.amazonaws.com/default/leetcode-stats?username=KnlnKS&theme=dark" width="400"/>
</a>

## Classic Cards
If you want to use the classic cards visit the legacy api at:  
https://github.com/KnlnKS/leetcode-stats/tree/master/vercel

## GitHub Actions
If you don't want to rely on an external api for your LeetCode stats I have provided a solution via GitHub Actions.
You can find it on the following branch: https://github.com/KnlnKS/leetcode-stats/tree/github-actions
