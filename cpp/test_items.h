/*
 * config class generated by excelizor tool.
 */
#include "../rapidjson/document.h"
#include "../rapidjson/istreamwrapper.h"
#include "../jsonconfig.h"
#include <stdio.h>
#include <cstdio>
#include <string>
#include <fstream>
#include <iostream>

using namespace std;
using namespace rapidjson;
namespace Configs
{
    class TestItems : public JsonConfig
    {
    public:
        int Id;
        
        string Name;
        
        int EffectType;
        
        float Price1;
        
        vector<int> Effect;
        
        vector<string> Introduce;
        
        bool IsUse;
        
    public:
        TestItems()
        {
            addMap("int", "Id", (char*)(&this->Id)-(char*)(this));
			addMap("string", "Name", (char*)(&this->Name)-(char*)(this));
			addMap("int", "EffectType", (char*)(&this->EffectType)-(char*)(this));
			addMap("float", "Price1", (char*)(&this->Price1)-(char*)(this));
			addMap("vector<int>", "Effect", (char*)(&this->Effect)-(char*)(this));
			addMap("vector<string>", "Introduce", (char*)(&this->Introduce)-(char*)(this));
			addMap("bool", "IsUse", (char*)(&this->IsUse)-(char*)(this));
        }
        ~TestItems() {}
    };

}//end namespace Configs