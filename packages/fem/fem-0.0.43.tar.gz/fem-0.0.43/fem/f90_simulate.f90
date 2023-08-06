subroutine simulate_data(par, nm, pq_sum, degs, p, q, k, x_idx, pdegs_sum, n, m, p_sum, l, x)
  implicit none

  ! par: m*n by sum(pq)
  !
  ! [ par(:, ipar(1)+1:ipar(2)) -- mn by pq(1) | par(:, ipar(2)+1:ipar(3)) -- mn by pq(2) | ... ]

  ! s: sum(p) by l
  !
  ! ------------------------------------
  ! s1: s(is(1)+1:is(2), :) -- p(1) by l
  ! ------------------------------------
  ! s2: s(is(2)+1:is(3), :) -- p(2) by l
  ! ------------------------------------
  ! ...
  ! ------------------------------------

  ! real*8 par(:,:)
  ! integer x_idx(:), degs(:), q(:), p(:)

  integer nm, pq_sum
  real*8 par(nm, pq_sum)
  ! f2py intent(hide), depend(par) :: nm=shape(par,0)
  ! f2py intent(hide), depend(par) :: pq_sum=shape(par,1)
  integer k
  integer degs(k)
  integer p(k)
  integer q(k)
  ! f2py intent(hide), depend(degs) :: k=len(degs)
  integer pdegs_sum
  integer x_idx(pdegs_sum)
  ! f2py intent(hide), depend(x_idx) :: pdegs_sum=len(x_idx)
  integer n
  integer m
  integer p_sum
  integer l
  integer, intent(out) :: x(p_sum,l)

  integer deg, pq(k), i_par(k+1), i_x(k+1), i_x_idx(k+1)
  integer it, id, iv, iv1, iv2, is, jt, jv, s
  real*8 wrk(nm), randr(n, l)

  call random_seed()
  call random_number(randr)

  pq = p*q

  i_x(1) = 0
  i_par(1) = 0
  i_x_idx(1) = 0
  i_x(2:) = p
  i_par(2:) = pq
  i_x_idx(2:) = p*degs
  do id=1,k
     i_x(id+1) = i_x(id+1) + i_x(id)
     i_par(id+1) = i_par(id+1) + i_par(id)
     i_x_idx(id+1) = i_x_idx(id+1) + i_x_idx(id)
  end do

  ! random initial condition
  do iv=1,n
     x(iv,1) = 1 + mod(int((10*m+1)*randr(iv,1)), m)
  end do
  ! fill in higher order (deg > 1) states
  it = 1
  do id=1,k
     deg = degs(id)
     if (deg==1) continue
     do iv=1,p(id)
        jv = (iv-1)*deg
        s = 0
        do is=1,deg
           s = s + (x(x_idx(i_x_idx(id) + jv + is), it)-1)*m**(is-1)
        end do
        x(i_x(id) + iv, it) = s + 1
     end do
  end do

  do it=2,l ! time
     ! compute local field
     wrk = 0
     jt = it-1
     do id=1,k ! degree
        do iv=1,p(id) ! variable
           ! print*, 'col', ipar(i2)+(i3-1)*q(i2)+s(is(i2)+i3,i1-1), 'between (inclusive)', ipar(i2)+1, ipar(i2+1)
           wrk = wrk + par(:,i_par(id)+(iv-1)*q(id)+x(i_x(id)+iv, jt))
        end do
     end do
     ! sample from Boltzmann distribution
     wrk = exp(wrk)
     do iv=1,n
        iv1 = (iv-1)*m+1; iv2 = iv*m
        wrk(iv1:iv2) = wrk(iv1:iv2) / sum(wrk(iv1:iv2))
        do is=1,m-1
           wrk(iv1+is) = wrk(iv1+is) + wrk(iv1+is-1)
        end do
        is=1
        do while(wrk(iv1+is-1) < randr(iv,it))
           is = is + 1
        end do
        x(iv,it) = is
     end do

     ! fill in higher order (deg > 1) states
     do id=1,k
        deg = degs(id)
        if (deg==1) continue
        do iv=1,p(id)
           jv = (iv-1)*deg
           s = 0
           do is=1,deg
              s = s + (x(x_idx(i_x_idx(id) + jv + is), it)-1)*m**(is-1)
           end do
           x(i_x(id) + iv, it) = s + 1
        end do
     end do

  end do

end subroutine simulate_data

! subroutine power_s(i1, m, k, degs, p, pdegs_sum, sidx, isidx, is, s)
!   implicit none

!   integer m, k, pdegs_sum
!   ! integer degs(k), p(k), sidx(pdegs_sum), isidx(k+1), is(k+1), s(p_sum,l)
!   integer degs(:), p(:), sidx(:), isidx(:), is(:), s(:,:)
!   integer i1, i2, i3, i4, deg, state

!   do i2=1,k
!      deg = degs(i2)
!      if (deg==1) continue
!      do i3=1,p(i2)
!         state = 0
!         do i4=1,deg
!            state = state + (s(sidx(isidx(i2) + (i3-1)*deg + i4), i1)-1)*m**(i4-1)
!         end do
!         s(is(i2) + i3, i1) = state + 1
!      end do
!   end do
! end subroutine power_s
