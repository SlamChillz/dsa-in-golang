package singly

/**
 * FindLoopNode: finds the loop node
 * @head: pointer to the head node of a linked list
 * Returns: returns the loop node
*/
func FindLoopNode(head *List) *List {
	if head == nil || head.next == nil {
		return nil
	}
	oneStep, twoSteps := head.next, head.next.next
	for twoSteps != nil && twoSteps.next != nil {
		if oneStep == twoSteps {
			break
		}
		oneStep = oneStep.next
		twoSteps = twoSteps.next.next
	}
	if oneStep != twoSteps {
		return nil
	}
	oneStep = head
	for oneStep != twoSteps {
		oneStep = oneStep.next
		twoSteps = twoSteps.next
	}
	return twoSteps
}