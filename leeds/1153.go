class Solution {
public:
    bool canConvert(string s, string t) {
        if (s == t) return true;
        int n = s.size();
        vector<int> to(26, -1);
        for (int i = 0; i < n; ++ i)
        {
            int x = s[i]-'a', y = t[i]-'a';
            if (to[x] == -1) to[x] = y;
            else if (to[x] != y) return false;
        }
        
        int has = 0;
        for (int i = 0; i < 26; ++ i)
            has += to[i] != -1;
        
        int flag = 0;
        for (int i = 0; i < 26; ++ i)
            for (int j = i+1; j < 26; ++ j)
                if (to[i] != -1 && to[j] != -1 && to[i] == to[j])
                {
                    flag = 1;
                }
        
        if (has != 26) return true;
        
        if (flag) return true;
        return false;
    }
};
