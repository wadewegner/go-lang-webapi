source params
rg=$1

ema.exe setup $rg -s $sub
ema.exe install $rg $ms -s $sub
ema.exe auth.service $rg $ms PublicAnon -s $sub