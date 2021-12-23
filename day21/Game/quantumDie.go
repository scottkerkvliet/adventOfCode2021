package game

// ************************************** Constants **************************************
const maxQuantumRoll = 3
const minQuantumRoll = 1
const maxQuantumTurn = rollsPerTurn * maxQuantumRoll
const minQuantumTurn = rollsPerTurn * minQuantumRoll

var universesPerTurn = map[int]int{}

func iterateSlice(slice []int) bool {
	i := 0
	for i < len(slice) {
		slice[i]++
		if slice[i] <= maxQuantumRoll {
			break
		}
		slice[i] = minQuantumRoll
		i++
	}

	return i < len(slice)
}

func init() {
	rolls := make([]int, rollsPerTurn)
	for i := 0; i < rollsPerTurn; i++ {
		rolls[i] = minQuantumRoll
	}
	rolls[0]--
	for iterateSlice(rolls) {
		sum := rolls[0]
		for i := 1; i < rollsPerTurn; i++ {
			sum += rolls[i]
		}
		universesPerTurn[sum] = universesPerTurn[sum] + 1
	}
}

// ************************************ Data structure ************************************
type QuantumDie struct {
	state      int
	isComplete bool
}

func NewQuantumDie() *QuantumDie {
	return &QuantumDie{minQuantumTurn, false}
}

func (d *QuantumDie) GetDie() (Die, int) {
	die := NewNumberDie(d.state)
	return die, universesPerTurn[d.state]
}

func (d *QuantumDie) IsComplete() bool {
	return d.isComplete
}

func (d *QuantumDie) IncrementState() {
	d.state++
	d.isComplete = d.state > maxQuantumTurn
}

func (d *QuantumDie) SetState(state int) {
	d.state = state
	d.isComplete = d.state > maxQuantumTurn
}

func (d *QuantumDie) GetState() int {
	return d.state
}
