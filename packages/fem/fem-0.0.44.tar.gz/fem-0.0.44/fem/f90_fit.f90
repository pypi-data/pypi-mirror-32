module fit_module
  use omp_lib
  implicit none

contains

  subroutine fit(x, y, n_x, n_y, m_x, m_y, m_y_sum, l, degs, n_deg, &
       x_oh_pinv1, x_oh_pinv2, x_oh_pinv3, x_oh_rank, n_s, &
       iters, overfit, par, disc, iter)
    implicit none

    integer, intent(in) :: x(n_x,l), y(n_y,l), m_x(n_x), m_y(n_y), degs(n_deg)
    integer, intent(in) :: n_x, n_y, m_y_sum, l, n_deg
    integer, intent(in) :: x_oh_rank, n_s
    real*8, intent(in) :: x_oh_pinv1(l, x_oh_rank)
    real*8, intent(in) :: x_oh_pinv2(x_oh_rank)
    real*8, intent(in) :: x_oh_pinv3(x_oh_rank, n_s)
    integer, intent(in) :: iters
    logical, intent(in) :: overfit
    !f2py intent(hide), depend(x) :: n_x=shape(x,0)
    !f2py intent(hide), depend(y) :: n_y=shape(y,0)
    !f2py intent(hide), depend(x) :: l=shape(x,1)
    !f2py intent(hide), depend(degs) :: n_deg=len(degs)
    !f2py intent(hide), depend(x_oh_pinv2) :: x_oh_rank=len(x_oh_pinv2)
    !f2py intent(hide), depend(x_oh_pinv3) :: n_s=shape(x_oh_pinv3,1)

    real*8, intent(out) :: par(m_y_sum, n_s)
    real*8, intent(out) :: disc(n_y)
    integer, intent(out) :: iter(n_y)

    integer i, ii, iii, j, k
    integer max_deg, deg, n_idx
    integer, allocatable :: bc(:)
    integer, allocatable :: n_idx_by_deg(:)
    integer, allocatable :: idx_ptr(:)
    integer, allocatable :: idx(:)
    integer, allocatable :: var(:)
    integer, allocatable :: stratifier(:)
    integer, allocatable :: s(:,:)
    integer, allocatable :: m_y_cumsum(:)

    ! print*, 'x_oh_pinv1', shape(x_oh_pinv1)
    ! print*, 'x_oh_pinv2', shape(x_oh_pinv2)
    ! print*, 'x_oh_pinv3', shape(x_oh_pinv3)
    ! print*, 'n_x', n_x
    ! print*, 'l', l

    ! print*, 'x'
    ! do i = 1, n_x
    !    print*, x(i, :5)
    ! end do

    max_deg = degs(n_deg)
    ! print*, 'n_deg', n_deg
    ! print*, 'max_deg', max_deg

    ! binomial coeff
    allocate(bc(max_deg))
    bc(1) = n_x
    bc(2:) = (/ (bc(i-1)*(n_x-i+1)/i, i=2, max_deg) /)
    ! print*, 'bc', bc

    allocate(n_idx_by_deg(n_deg))
    n_idx_by_deg = (/ (bc(degs(i)), i=1, n_deg) /)
    n_idx = sum(n_idx_by_deg)
    ! print*, 'n_idx_by_deg', n_idx_by_deg
    ! print*, 'n_idx', n_idx

    allocate(idx_ptr(n_idx+1))
    idx_ptr(1) = 0
    i = 2
    do j = 1, n_deg
       do k = 1, n_idx_by_deg(degs(j))
          idx_ptr(i) = idx_ptr(i-1) + degs(j)
          i = i + 1
       end do
    end do
    ! print*, 'idx_ptr', idx_ptr

    allocate(idx(sum(n_idx_by_deg*degs)))
    i = 1
    do j = 1, n_deg
       allocate(var(degs(j)))
       var = (/ (i, i=1, degs(j)) /)
       idx(idx_ptr(i)+1:idx_ptr(i+1)) = var
       i = i + 1
       do k = 1, n_idx_by_deg(degs(j))-1
          ii = degs(j)
          var(ii) = var(ii) + 1
          do while (var(ii) + degs(j) > ii + n_x)
             ii = ii - 1
             var(ii) = var(ii) + 1
          end do
          var(ii+1:) = (/ (var(iii-1)+1, iii=ii+1, degs(j)) /)
          idx(idx_ptr(i)+1:idx_ptr(i+1)) = var
          i = i + 1
       end do
       deallocate(var)
    end do
    ! print*, 'idx', idx

    allocate(stratifier(n_idx))
    stratifier = 1
    do i = 2, n_idx
       do j = idx_ptr(i-1)+1, idx_ptr(i)
          stratifier(i) = stratifier(i) * m_x(idx(j))
       end do
       stratifier(i) = stratifier(i) + stratifier(i-1)
    end do
    ! print*, 'stratifier', stratifier

    ! powers of x stratified, i.e. cols of par
    allocate(s(n_idx, l))
    do i = 1, n_idx
       s(i, :) = x(idx(idx_ptr(i)+1), :)
       do j = idx_ptr(i)+2, idx_ptr(i+1)
          s(i, :) = s(i, :)*m_x(idx(j)) + x(idx(j), :)
       end do
       s(i, :) = s(i, :) + stratifier(i)
    end do
    ! print*, 's'
    ! do i = 1, n_idx
    !    print*, s(i, :5)
    ! end do

    allocate(m_y_cumsum(n_y+1))
    m_y_cumsum(1) = 0
    do i = 1, n_y
       m_y_cumsum(i+1) = m_y_cumsum(i) + m_y(i)
    end do

    ! print*, 'm_y_cumsum', m_y_cumsum
    ! print*, 'n_s', n_s
    ! print*, 'par', shape(par)

    !$omp parallel do
    do i = 1, n_y
       call fit_i(s, n_idx, y(i,:)+1, m_y(i), l, &
            x_oh_pinv1, x_oh_pinv2, x_oh_pinv3, x_oh_rank, n_s, &
            iters, overfit, par(m_y_cumsum(i)+1:m_y_cumsum(i+1),:), disc(i), iter(i))
      ! print*, i, iter(i), disc(i), m_y_cumsum(i)+1, m_y_cumsum(i+1)
    end do
    !$omp end parallel do

    deallocate(bc, n_idx_by_deg, idx_ptr, idx, stratifier, s, m_y_cumsum)

  end subroutine fit

  subroutine fit_i(s, n_idx, y, m_y, l, &
       x_oh_pinv1, x_oh_pinv2, x_oh_pinv3, x_oh_rank, n_s, &
       iters, overfit, par, disc, iter)
    implicit none

    integer, intent(in) :: n_idx, m_y, l, x_oh_rank, n_s, iters
    integer, intent(in) :: s(n_idx, l), y(l)
    real*8, intent(in) :: x_oh_pinv1(l, x_oh_rank)
    real*8, intent(in) :: x_oh_pinv2(x_oh_rank)
    real*8, intent(in) :: x_oh_pinv3(x_oh_rank, n_s)
    logical, intent(in) :: overfit

    real*8, intent(out) :: par(m_y, n_s)
    real*8, intent(out) :: disc
    integer, intent(out) :: iter

    real*8 last_disc
    logical disc_mask(m_y, l)
    integer n_wrong_states
    real*8 wrk(m_y, l)
    real*8 dpar(m_y, x_oh_rank)
    integer i, t

    par = 0
    disc = 1.0 / m_y / m_y + 1.0
    n_wrong_states = (m_y - 1) * l

    if (.not.overfit) then
       disc_mask = .true.
       do t = 1, l
          disc_mask(y(t), t) = .false.
       end do
    end if

    do iter = 2, iters-1

       ! compute energies
       wrk = 0
       do t = 1, l
          do i = 1, n_idx
             wrk(:, t) = wrk(:, t) + par(:, s(i, t))
          end do
       end do

       ! probabilities
       wrk = exp(wrk)
       do t = 1, l
          wrk(:, t) = wrk(:, t) / sum(wrk(:, t))
       end do

       ! discrepancy
       if (.not.overfit) then
          last_disc = disc
          disc = sum(wrk*wrk, disc_mask) / n_wrong_states
          if (disc > last_disc) then
             exit
          end if
       end if

       do t = 1, l
          wrk(y(t), t) = wrk(y(t), t) - 1
       end do

       dpar = matmul(wrk, x_oh_pinv1)
       do i = 1, x_oh_rank
          dpar(:, i) = dpar(:,i) * x_oh_pinv2(i)
       end do

       par = par - matmul(dpar, x_oh_pinv3)

    end do

    disc = sum(wrk*wrk, disc_mask) / n_wrong_states

  end subroutine fit_i

end module fit_module
