class Solution {
public:
    int maxRepOpt1(string text) {
        int n=text.size(),i,j,k,ans=0,s,a[20005],t[20005];
        char c;
        a[0]=0;
        for(c='a';c<='z';c++)
        {
            a[0]=0;
            for(i=0;i<n;i++)a[i+1]=text[i]==c;
            for(i=1;i<=n;i++)a[i]+=a[i-1];
            if(!a[n])continue;
            for(i=0;i<=n;i++)a[i]+=n-i;
            fill(t+1,t+n+2,n+1);
            for(i=s=0;i<=n;i++)
            {
                for(j=a[i]+1,k=i;j;j^=j&-j)k=min(k,t[j]);
                s=max(s,i-k);
                for(j=a[i];j<=n+1;j+=j&-j)t[j]=min(t[j],i);
            }
            ans=max(ans,min(s,a[n]));
        }
        return ans;
    }
};

