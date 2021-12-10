#include <iostream>
#include <list>
#include <vector>
using namespace std;

list <int> * mylist;
/*int main() {
    mylist->assign(2, 3);
    for (list<int>::iterator i = mylist->begin(); i != mylist->end(); i++) {
        cout << *i << " ";
    }
}*/

//___________________________________________
int car(vector<int> * v) {
    return v->begin();
}

vector<int> cdr(vector<int> * v) {
    vector <int> * nvec;
    int k = 0;
    for (vector<int>::iterator i = advance(v->begin(),1); i != v->end(); i++) {
        nvec->insert(k, 1, i);
        k++;
    }
    return nvec;
}

vector<int> cons(int e, vector<int> * v) {
    vector <int> * nvec;
    nvec->insert(0, 1, e);
    int k = 1;
    for (vector<int>::iterator i = v->begin(); i != v->end(); i++) {
        nvec->insert(k, 1, i);
        k++;
    }
    return nvec;
}

int main() {
    vector <int> * myvec;
    myvec->insert(0, 1, 1);
    myvec->insert(1, 1, 2);
    cout << car(myvec) << "\n";
    cdr(myvec);
    cons(0, myvec);
}