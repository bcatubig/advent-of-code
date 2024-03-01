#ifndef DAY_1_H
#define DAY_1_H

#include <stdio.h>

struct ListElement {
	int data;
	struct ListElement *next;
};

struct List {
	int size;
	struct ListElement *head;
	struct ListElement *tail;
};

void list_init(struct List *list);

void list_destroy(struct List *list);

int (*match)(const void *key1, const void *key2);

int list_ins_next(struct List *list, struct ListElement *element, int data);
int list_rem_next(struct List *list, struct ListElement *element);

#define list_size(list) ((list)->size)
#define list_head(list) ((list)->head)
#define list_tail(list) ((list)->tail)
#define list_is_head(list, element) ((element) == (list)->head ? 1 : 0)
#define list_is_tail(element) ((element)->next == NULL ? 1 : 0)
#define list_next(element) ((element)->next)
#define list_data(element) ((element)->data)

/*
* Function: read_calibration_value
* Computes the calibration value for launching a trebuchet
*
* f: the calibration document
*
* returns: the calibration value. -1 if an error occured
*/
int read_calibration_value(FILE *f);

#endif
