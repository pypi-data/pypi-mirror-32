subroutine fit(x, y, p_sum, l, ohx_pinv1, ohx_pinv2, ohx_pinv3, ohx_rank, pq_sum, q, p, k, m_y, max_iters, iters, iter, par, disc)
  implicit none

  ! x(sum(p), l)
  ! y(l)
  ! ohx_pinv1(l, ohx_rank)
  ! ohx_pinv2(ohx_rank)
  ! ohx_pinv3(ohx_rank, pq_sum)
  ! p(k+1), q(k+1)

  integer p_sum
  integer l
  integer x(p_sum, l)
  integer y(l)
  ! f2py intent(hide), depend(x) :: p_sum=shape(x,0)
  ! f2py intent(hide), depend(x) :: l=shape(x,1)
  integer ohx_rank
  integer pq_sum
  real*8 ohx_pinv1(l, ohx_rank)
  real*8 ohx_pinv2(ohx_rank)
  real*8 ohx_pinv3(ohx_rank, pq_sum)
  ! f2py intent(hide), depend(ohx_pinv1) :: ohx_rank=shape(ohx_pinv1,1)
  ! f2py intent(hide), depend(ohx_pinv3) :: ohx_rank=shape(ohx_pinv3,1)
  integer k
  integer p(k)
  integer q(k)
  ! f2py intent(hide), depend(p) :: k=len(p)
  integer m_y
  integer max_iters
  integer, optional :: iters
  integer, intent(out) :: iter
  real*8, intent(out) :: par(max_iters, m_y, pq_sum)
  ! f2py depend(max_iters) :: par
  ! f2py depend(m_y) :: par
  ! f2py depend(pq_sum) :: par
  real*8, intent(out) :: disc(max_iters)
  ! f2py depend(disc) :: max_iters
  logical disc_stop

  real*8 wrk(m_y, l), dpar(m_y, ohx_rank)
  integer pq(k), i_par(k+1), i_x(k+1)
  integer it, id, iv, j, n_wrong_states
  logical mask(m_y, l)

  if (present(iters)) then
     disc_stop = .false.
  else
     disc_stop = .true.
     iters = max_iters
  end if

  n_wrong_states = l*(m_y-1)

  pq = p*q

  i_x(1) = 0
  i_par(1) = 0
  i_x(2:) = p
  i_par(2:) = pq
  do id=1,k
     i_x(id+1) = i_x(id+1) + i_x(id)
     i_par(id+1) = i_par(id+1) + i_par(id)
  end do

  mask = .true.
  do it=1,l
     mask(y(it),it) = .false.
  end do

  iter = 1
  par = 0
  disc(1) = 1.0/m_y/m_y + 1

  do iter=2, iters

     ! compute energies
     wrk = 0
     do it=1,l ! time
        do id=1,k ! degree
           do iv=1,p(id) ! variable
              wrk(:,it) = wrk(:,it) + par(iter-1,:,i_par(id)+(iv-1)*q(id)+x(i_x(id)+iv,it))
           end do
        end do
     end do

     ! compute probabilities
     wrk = exp(wrk)
     do it=1,l
        wrk(:,it) = wrk(:,it) / sum(wrk(:,it))
     end do

     ! discrepancy
     disc(iter) = sum(wrk*wrk, mask) / n_wrong_states

     if (disc_stop.and.(disc(iter) > disc(iter-1))) then
        exit
     end if

     do it=1,l
        wrk(y(it),it) = wrk(y(it),it) - 1
     end do

     dpar = matmul(wrk, ohx_pinv1)
     do j=1,ohx_rank
        dpar(:,j) = dpar(:,j)*ohx_pinv2(j)
     end do
     par(iter,:,:) = par(iter-1,:,:) - matmul(dpar, ohx_pinv3)

  end do

end subroutine fit
