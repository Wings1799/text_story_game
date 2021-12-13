#include <iostream>
#include <string>
#include <list>
#include <vector>
using namespace std;

int main() {
    list<int> mylist;
    mylist.assign(1, 1);
    mylist.assign(3, 2);
    int k = 3;
    list<int>::iterator i;
    for (i = mylist.begin(); i != mylist.end(); i++) {
        mylist.insert(i, k, 99);
    }
    vector<string> paragraph;
    ifstream file ("test.txt");
    if (file) {
        copy (istream_iterator<string> (file), istream_iterator<string>(), back_inserter (paragraph));
    }
}