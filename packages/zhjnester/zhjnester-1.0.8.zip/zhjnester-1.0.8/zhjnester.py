"""
����һ����ʾǶ�׷�����ʾ�������ܴ�ӡ��list���͵����ݵ�Ԫ����
"""
__author__ = 'hjzheng'
#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
��ӡ��list���͵����ݵ�Ԫ����
"""
def print_lol(the_list,indent=False,level=0):
   for each_item in the_list:
        if isinstance(each_item,list):
            print_lol(each_item,indent,level+1)
        else:
            if indent:
                for tab_stop in range(level):
                    print("\t"),
            print(each_item)

"""
import zhjnester
movies=['�����',1975,'����ı',91,['����',['̫���ճ�����','���ӵ���','��ӹ�']]]
zhjnester.print_lol(movies,4)
"""