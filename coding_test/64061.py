def solution(board, moves):
    answer = 0
    stack = []
    
    # 보드의 몇번째 칸까지 인형이 있는지 확인
    board_chk = [-1 for _ in range(len(board))]

    for i in moves:
        # 해당 열에 인형이 없으면 넘어감
        if board_chk[i - 1] == len(board) + 1:
            continue
        # 해당 열을 체크 한 적이 없을 경우
        elif board_chk[i - 1] == -1:
            for j in range(len(board)):
                # 0이 아닌 값을 발견한 경우
                if board[j][i - 1]:
                    board_chk[i - 1] = j + 1
                    stack.append(board[board_chk[i - 1] - 1][i - 1])
                    board[board_chk[i - 1] - 1][i - 1] = 0
                    board_chk[i - 1] += 1
                    break
                # 해당 열에서 인형을 발견 못한 경우 
                elif j == len(board) - 1:
                    board_chk[i - 1] = len(board)
        # 해당 열을 체크한 적 있는 경우
        else:
            # 해당 칸에 인형이 있는 경우, 전부 다 비어 있는 경우 0 이 들어감을 방지
            if board[board_chk[i - 1] - 1]:
                stack.append(board[board_chk[i - 1] - 1][i - 1])
                board[board_chk[i - 1] - 1][i - 1] = 0
                board_chk[i - 1] += 1
        # 스택에 2개 이상 있는 경우
        if len(stack) >= 2:
            # 마지막 두개의 값이 같으면
            if stack[-1] == stack[-2]:
                stack.pop()
                stack.pop()
                answer += 2
            
    return answer


print(solution([[0,0,0,0,0],[0,0,1,0,3],[0,2,5,0,1],[4,2,4,4,2],[3,5,1,3,1]], 
          [1,5,3,5,1,2,1,4]))