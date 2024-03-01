#include <string.h>
#include <stdlib.h>
#include <stdio.h>
#include <ctype.h>
#include "day_1.h"

void list_init(struct List *list)
{
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

int list_ins_tail(struct List *list, int data)
{
	int rc = list_ins_next(list, list->tail, data);

	if (rc != 0) {
		return -1;
	}

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
	int result = 0;

	// For line in file, read each line value, add to result

	return result;
}

int read_line_value(const char *line)
{
	struct List *nums;

	nums = (struct List *)malloc(sizeof(struct List));

	if (nums == NULL) {
		return -1;
	}

	list_init(nums);

	size_t length = strlen(line);
	for (int i = 0; i < length; i++) {
		if (isdigit(line[i])) {
			int num = atoi(&line[i]);
			if ((list_ins_tail(nums, num)) != 0) {
				fprintf(stderr,
					"failed to add %d to end of list\n",
					num);
				continue;
			}
		}
	}

	int result = 0;

	if (list_size(nums) == 1) {
		char buf[2];
		int num = nums->head->data;

		buf[0] = num + '0';
		buf[1] = num + '0';

		result = atoi(buf);

	} else {
		char buf[2];

		buf[0] = nums->head->data + '0';
		buf[1] = nums->tail->data + '0';

		result = atoi(buf);
	}

	list_destroy(nums);

	return result;
}
