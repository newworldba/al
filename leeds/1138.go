class Solution {
public:
    string alphabetBoardPath(string t) {
        int n=t.size(),pos[27][2],i,j;
        string ans;
        string s[6]={"abcde", "fghij", "klmno", "pqrst", "uvwxy", "z****"};
        for (i=0;i<6;i++)
        {for (j=0;j<5;j++)
        {if (s[i][j]>='a'&&s[i][j]<='z')
        {pos[s[i][j]-'a'][0]=i;pos[s[i][j]-'a'][1]=j;}
        }
        }
        int nx=0,ny=0;
        for (i=0;i<n;i++)
        {int o=t[i]-'a';
         if (nx==5&&ny==0)
         {while (nx>pos[o][0]) {ans+="U";nx--;}
         while (nx<pos[o][0]) {ans+="D";nx++;}
         while (ny>pos[o][1]) {ans+="L";ny--;}
         while (ny<pos[o][1]) {ans+="R";ny++;}
         }
         else
         {while (ny>pos[o][1]) {ans+="L";ny--;}
         while (ny<pos[o][1]) {ans+="R";ny++;}
         while (nx>pos[o][0]) {ans+="U";nx--;}
         while (nx<pos[o][0]) {ans+="D";nx++;}
         }
         
          ans+="!";
         
        }
        return ans;
        
    }
};
