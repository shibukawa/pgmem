package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SnapBuildProcessChange(m *base.Module, l0 int32, l1 int32, l2 int64) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int64
	_ = v62
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	v4 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v7 <= v4 {
		v90 = v4
		return v90
	} else {
		if v7 == int32(1) {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
			if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v12))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(l1)) == int32(0) {
				v24 = base.B2i32(base.Ui32(l1) < base.Ui32(v12))
			} else {
				v24 = int32(base.Ui32(l1-v12) >> (uint(int32(31)) % 32))
			}
			if v24 != 0 {
				v90 = v4
				return v90
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
				v27 = F_ReorderBufferXidHasBaseSnapshot(m, v26, l1)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					if v27 != 0 {
						v90 = int32(1)
						return v90
					} else {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if v31 == int32(0) {
							v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
							v40 = F_MemoryContextAllocZero(m, v34, v35<<(uint(int32(2))%32)+int32(76))
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v40))) = int32(5)
								v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v40)+4)) = v44
								v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v48 = v40 + int32(72)
								*(*int32)(unsafe.Add(mBase, uint32(v40)+12)) = v48
								*(*int32)(unsafe.Add(mBase, uint32(v40)+8)) = v46
								v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
								*(*int32)(unsafe.Add(mBase, uint32(v40)+16)) = v51
								v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
								v55 = v51 << (uint(int32(2)) % 32)
								if v55 != 0 {
									v56 = F__emscripten_memcpy_bulkmem(m, v48, v53, v55)
									mBase = m.M
									v57 = v56
								} else {
									v57 = v48
								}
								F_pg_qsort(m, v57, v51, int32(4), int32(185))
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return int32(0)
								} else {
									v62 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v40)+64)) = v62
									*(*int64)(unsafe.Add(mBase, uint32(v40)+44)) = v62
									v66 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v66
									*(*int64)(unsafe.Add(mBase, uint32(v40)+20)) = v62
									*(*int32)(unsafe.Add(mBase, uint32(v40)+27)) = v66
									*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v40
									v73 = *(*int32)(unsafe.Add(mBase, uint32(v40)+44))
									*(*int32)(unsafe.Add(mBase, uint32(v40)+44)) = v73 + int32(1)
									v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
									v79 = v77
									v81 = int32(1)
									v82 = *(*int32)(unsafe.Add(mBase, uint32(v79)+44))
									*(*int32)(unsafe.Add(mBase, uint32(v79)+44)) = v82 + v81
									v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
									v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
									F_ReorderBufferSetBaseSnapshot(m, v86, l1, l2, v87)
									mBase = m.M
									v89 = m.ExcPending
									if v89 != 0 {
										return int32(0)
									} else {
										v90 = v81
										return v90
									}
								}
							}
						} else {
							v79 = v31
							v81 = int32(1)
							v82 = *(*int32)(unsafe.Add(mBase, uint32(v79)+44))
							*(*int32)(unsafe.Add(mBase, uint32(v79)+44)) = v82 + v81
							v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
							v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							F_ReorderBufferSetBaseSnapshot(m, v86, l1, l2, v87)
							mBase = m.M
							v89 = m.ExcPending
							if v89 != 0 {
								return int32(0)
							} else {
								v90 = v81
								return v90
							}
						}
					}
				}
			}
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
			v27 = F_ReorderBufferXidHasBaseSnapshot(m, v26, l1)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				if v27 != 0 {
					v90 = int32(1)
					return v90
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v31 == int32(0) {
						v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
						v40 = F_MemoryContextAllocZero(m, v34, v35<<(uint(int32(2))%32)+int32(76))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v40))) = int32(5)
							v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v40)+4)) = v44
							v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v48 = v40 + int32(72)
							*(*int32)(unsafe.Add(mBase, uint32(v40)+12)) = v48
							*(*int32)(unsafe.Add(mBase, uint32(v40)+8)) = v46
							v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
							*(*int32)(unsafe.Add(mBase, uint32(v40)+16)) = v51
							v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
							v55 = v51 << (uint(int32(2)) % 32)
							if v55 != 0 {
								v56 = F__emscripten_memcpy_bulkmem(m, v48, v53, v55)
								mBase = m.M
								v57 = v56
							} else {
								v57 = v48
							}
							F_pg_qsort(m, v57, v51, int32(4), int32(185))
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								v62 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v40)+64)) = v62
								*(*int64)(unsafe.Add(mBase, uint32(v40)+44)) = v62
								v66 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v66
								*(*int64)(unsafe.Add(mBase, uint32(v40)+20)) = v62
								*(*int32)(unsafe.Add(mBase, uint32(v40)+27)) = v66
								*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v40
								v73 = *(*int32)(unsafe.Add(mBase, uint32(v40)+44))
								*(*int32)(unsafe.Add(mBase, uint32(v40)+44)) = v73 + int32(1)
								v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
								v79 = v77
								v81 = int32(1)
								v82 = *(*int32)(unsafe.Add(mBase, uint32(v79)+44))
								*(*int32)(unsafe.Add(mBase, uint32(v79)+44)) = v82 + v81
								v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
								v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
								F_ReorderBufferSetBaseSnapshot(m, v86, l1, l2, v87)
								mBase = m.M
								v89 = m.ExcPending
								if v89 != 0 {
									return int32(0)
								} else {
									v90 = v81
									return v90
								}
							}
						}
					} else {
						v79 = v31
						v81 = int32(1)
						v82 = *(*int32)(unsafe.Add(mBase, uint32(v79)+44))
						*(*int32)(unsafe.Add(mBase, uint32(v79)+44)) = v82 + v81
						v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
						v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						F_ReorderBufferSetBaseSnapshot(m, v86, l1, l2, v87)
						mBase = m.M
						v89 = m.ExcPending
						if v89 != 0 {
							return int32(0)
						} else {
							v90 = v81
							return v90
						}
					}
				}
			}
		}
	}
}
