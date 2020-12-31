
#include <iostream>
#include <string.h>
#include "mq.h"

using namespace std;

void MqMessage( char *msg ) {
    cout << "msg: " << msg << endl;
}

int main()
{
    cout << "Hello World!" << endl;

    if ( Init() == false ) {
        cout << "Init 失败" << endl;
        return -1;
    }

    char* name = strdup("Center");
    char* msg  = strdup("HelloCenter");
    if ( Send( name, msg ) == false ) {
        cout << "Init 失败" << endl;
        return -1;
    }

    Recv( strdup( "KWX_TEST" ), MqMessage );

    Release();
    return 0;
}
