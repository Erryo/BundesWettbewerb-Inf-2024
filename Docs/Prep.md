# Requirements 
1. Take in values for: - Width - Length - Number of parties 
2. Ensure at least as many lots as parties. 
3. No more than 10% extra lots. 
4. Make the lots as square as possible.
# Solutions
## 1
- **Easiest**: Use `os.Args`.
- **Not Worth It**: Use `ReadLine`.
- **Hardest (Best Solution)**: Use a config file.
## 2. 
![[Drawing 2024-09-16 16.17.56.excalidraw]]

Create dividers for the 2 unequal sides
 - `a / DividerA` == `a/dA`
 - `b / DividerB` == `b/dB`
\

| a   | b   | dA  | dB  | a1  | b1  |
| --- | --- | --- | --- | --- | --- |
| 1   | 2   | 4   | 8   |     |     |


> [!NOTE]
> Dividing the bigger side by bigger numbers

> [!NOTE]
> S is the ratio of squareness: `S=a1/b1`

## 3.
At the beginning, calculate the 10% limit.
## 4. 
Ensure lots are as square as possible.

### Properties of a Square:

- All sides are equal.
- Area = \(a^2\).
- Perimeter = \(4a\).
- The ratio between two sides = 1.
![[Drawing 2024-09-16 16.17.56.excalidraw]]