class Solution {
public:
    int stoneGameII(vector<int>& a) {
        int dp[105][105],s[105],n=a.size(),i,j,k;
        memset(dp,0,sizeof(dp));
        memset(s,0,sizeof(s));
        for (i=1;i<=n;i++) s[i]=s[i-1]+a[i-1];
        for (i=n;i>=1;i--)
        {for (j=1;j<=100;j++)
        {for (k=1;i+k-1<=n&&k<=2*j;k++)
        {dp[i][j]=max(dp[i][j],s[n]-s[i-1]-dp[i+k][min(100,max(j,k))]);}
        }
        }
        return dp[1][1];
    }
};
