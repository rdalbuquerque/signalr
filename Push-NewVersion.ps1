param(
    [string]$CommitMsg = "versioning",
    [switch]$Alpha,
    [switch]$Beta
)

git add .
git commit -m $CommitMsg

if ($Alpha -and $Beta) {
    Write-Warning "You can't use both Alpha and Beta flags at the same time"
    exit
}

if ($Alpha) {
    standard-version --prerelease alpha
} elseif ($Beta) {
    standard-version --prerelease beta
} else {
    standard-version
}

git push --follow-tags origin master