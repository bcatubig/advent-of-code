#ifndef DAY_1_H
#define DAY_1_H

#include <stdio.h>

void say_hello(void);

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

int list_ins_next(struct List *list, struct ListElement *element, int data);
int list_rem_next(struct List *list, struct ListElement *element);
int list_ins_tail(struct List *list, int data);

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

/*
* Function: read_line_value
* Parses the calibration value for a given line
*
* line: line in file to parse
*
* returns: the calibration value for a given line. -1 if an error occured
*/
int read_line_value(const char *line);

#endif
