#include <ctype.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "day_1.h"

void list_init(struct List *list)
{
	memset(list, 0, sizeof(struct List));
	list->size = 0;
	list->head = NULL;
	list->tail = NULL;
}

void list_destroy(struct List *list)
{
	while (list_size(list) > 0) {
		list_rem_next(list, NULL);
	}
	memset(list, 0, sizeof(struct List));
}

int list_ins_next(struct List *list, struct ListElement *element, int data)
{
	struct ListElement *new_element;

	new_element = (struct ListElement *)malloc(sizeof(struct ListElement));

	if (new_element == NULL) {
		return -1;
	}

	new_element->data = data;

	// Handle insertion at head of list
	if (element == NULL) {
		if (list_size(list) == 0) {
			list->tail = new_element;
		}

		new_element->next = list->head;
		list->head = new_element;
	} else {
		if (element->next == NULL) {
			list->tail = new_element;
		}

		new_element->next = element->next;
		element->next = new_element;
	}

	list->size++;

	return 0;
}

int list_rem_next(struct List *list, struct ListElement *element)
{
	struct ListElement *old_element;

	if (list_size(list) == 0) {
		return -1;
	}

	// remove the head
	if (element == NULL) {
		old_element = list->head;
		list->head = list->head->next;

		if (list_size(list) == 1) {
			list->tail = NULL;
		}
	} else {
		if (element->next == NULL) {
			return -1;
		}

		old_element = element->next;
		if (element->next == NULL) {
			list->tail = element;
		}
	}

	free(old_element);

	list->size--;

	return 0;
}

int read_calibration_value(FILE *f)
{
	// thing
	return 0;
}

int read_line_value(char *line, size_t line_length)
{
	struct List *nums;

	nums = (struct List *)(sizeof(struct List));

	if (nums == NULL) {
		return -1;
	}

	for (int i = 0; i < line_length - 1; i++) {
		if isdigit (line[i]) {
			// add to list
			int rv;
		}
		printf("%c - isNum: %d\n", line[i], isdigit(line[i]) != 0);
	}

	return 0;
}
