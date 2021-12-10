#include <iostream>
#include <list>
#include <vector>
using namespace std;

list <int> * mylist;
int main() {
    mylist->assign(2, 3);
    for (list<int>::iterator i = mylist->begin(); i != mylist->end(); i++) {
        cout << *i << " ";
    }
}

//___________________________________________
int car(vector<int> v) {
    return v->begin();
}

vector<int> cdr(vector<int> v) {
    vector <int> * nvec;
    int k = 0;
    for (vector<int>::iterator i = advance(v->begin(),1); i != v->end(); i++) {
        nvec->insert(i, k);
        k++;
    }
    return nvec;
}

vector<int> cons(int e, vector<int> v) {
    vector <int> * nvec;
    nvec->insert(e, 0);
    int k = 1;
    for (vector<int>::iterator i = v->begin(); i != v->end(); i++) {
        nvec->insert(i, k);
        k++;
    }
    return nvec;
}