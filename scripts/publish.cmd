source params
nuget=$1

NuGet SetApiKey $apiKey -source "$source"
NuGet push $nuget -source "$source"