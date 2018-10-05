#!/bin/bash
#!Coded by Suleman Malik
#!www.sulemanmalik.com

cont(){
rm /tmp/cnames
rm /tmp/cname-ln
rm /tmp/fn-py
}
bl(){
	echo -e "\033[0;31m-------------------------------------------------------------\033[0m"
}
ban(){
echo '''
  ________                                                                   
 /__  ___/                                                           
   / /   ___     / ___      ___      ___              ___      __    
  / /  //   ) ) //\ \     //___) ) //   ) ) ||  / / //___) ) //  ) ) 
 / /  //   / / //  \ \   //       //   / /  || / / //       //       
/ /  ((___( ( //    \ \ ((____   ((___/ /   ||/ / ((____   //        
'''
}
cont > /dev/null 2>&1
clear                                                                        
ban
echo -e "\033[0;32m  @sulemanmalik_3                         		 v1\033[0m"
bl
echo "File Name:"
read inp
echo ''
echo ''
cot=0
lof=$(wc -l < $inp | sed 's/ //g')
while read dom
do 
cot=$(($cot + 1))
echo -ne "\033[0;32mScanning Subdomains:\033[0m $cot/$lof\r"
cname=$( dig CNAME $dom | grep "CNAME" | tail -n1 | cut -c"29-" | sed 's/^.*E//' >> /tmp/cnames)
echo "$cot - $dom --> " >> /tmp/cname-ln
done < $inp
prog=$(awk 'NR==FNR{a[++y]=$0;next}{b[++x]=$0}
END{z=x>y?x:y;while(++i<=z){print a[i],b[i]}}' /tmp/cname-ln /tmp/cnames > /tmp/fn-py)
echo ''
echo ''
cat /tmp/fn-py
sleep 1
echo ''
bl
cont > /dev/null 2>&1
