class Solution {
public:
    int largest1BorderedSquare(vector<vector<int>>& g) {
        int s1[105][105],s2[105][105],ans=0;
        memset(s1,0,sizeof(s1));
        memset(s2,0,sizeof(s2));
        int n=g.size(),m=g[0].size(),i,j,k;
        for (i=1;i<=n;i++)
        {for (j=1;j<=m;j++)
        {s1[i][j]=s1[i][j-1]+g[i-1][j-1];
         s2[i][j]=s2[i-1][j]+g[i-1][j-1];
        }
        }
        for (i=1;i<=n;i++)
        {for (j=1;j<=m;j++)
        {for (k=1;i+k-1<=n&&j+k-1<=m;k++)
        {int l1=i,r1=i+k-1;
         int l2=j,r2=j+k-1;
         if (k<=ans) continue;
         if (s1[l1][r2]-s1[l1][l2-1]!=k) continue;
         if (s1[r1][r2]-s1[r1][l2-1]!=k) continue;
         if (s2[r1][r2]-s2[l1-1][r2]!=k) continue;
         if (s2[r1][l2]-s2[l1-1][l2]!=k) continue;
         ans=k;
        }
        }
        }
        return ans*ans;
        
    }
};
