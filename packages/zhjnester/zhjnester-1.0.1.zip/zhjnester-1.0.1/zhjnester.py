"""
����һ����ʾǶ�׷�����ʾ�������ܴ�ӡ��list���͵����ݵ�Ԫ����
"""
__author__ = 'hjzheng'
#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
��ӡ��list���͵����ݵ�Ԫ����
"""
def print_lol(the_list,level):
   for each_item in the_list:
        if isinstance(each_item,list):
            print_lol(each_item,level)
        else:
            for tab_stop in range(level):
                print("\t")
            print(each_item)

