class Solution {
public:
    vector<string> mostVisitedPattern(vector<string>& user, vector<int>& time, vector<string>& web) {
        int N = user.size();
        
        set<string> webs;
        for (auto s : web) webs.insert(s);
        
        map<string, set<pair<int, string> > > v;
        for (int i = 0; i < N; ++ i)
            v[user[i]].insert(make_pair(time[i], web[i]));
        
        int ans = 0;
        vector<string> ret;
        for (auto a : webs)
            for (auto b : webs)
                for (auto c : webs)
                {
                    int num = 0;
                    for (auto &[user, history] : v)
                    {
                        int cnt = 0;
                        for (auto &[time, web] : history)
                        {
                            if (cnt == 0 && web == a) cnt ++;
                            else if (cnt == 1 && web == b) cnt ++;
                            else if (cnt == 2 && web == c)
                            {
                                cnt ++;
                                break;
                            }
                        }
                        if (cnt == 3) num ++;
                    }
                    if (num > ans)
                    {
                        ans = num;
                        ret = {a, b, c};
                    }
                }
        
        return ret;
    }
};
