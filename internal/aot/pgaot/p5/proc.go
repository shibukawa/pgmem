package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ProcArrayRemove(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v31 int32
	_ = v31
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v50 int64
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayRemove[0]))
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayRemove[1]))
	v18 = F_LWLockAcquire(m, v14+int32(512), int32(0))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayRemove[1]))
		v25 = F_LWLockAcquire(m, v21+int32(384), int32(0))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return
		} else {
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			if l1 != 0 {
				v29 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayRemove[2]))
				v30 = *(*int64)(unsafe.Add(mBase, uint32(v29)+48))
				v31 = base.I32_wrap_i64(v30)
				if base.B2i32(base.Ui32(int32(2)) < base.Ui32(l1))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v31)) == int32(0) {
					v43 = base.B2i32(base.Ui32(v31) < base.Ui32(l1))
				} else {
					v43 = int32(base.Ui32(v31-l1) >> (uint(int32(31)) % 32))
				}
				v45 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayRemove[2]))
				if v43 != 0 {
					*(*int64)(unsafe.Add(mBase, uint32(v45)+48)) = v30 + base.I64_extend_i32_s(l1-v31)
				} else {
				}
				v50 = *(*int64)(unsafe.Add(mBase, uint32(v45)+56))
				*(*int64)(unsafe.Add(mBase, uint32(v45)+56)) = v50 + int64(1)
				v54 = int32(_a_F_ProcArrayRemove_0)
				v55 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayRemove[3]))
				v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
				v60 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v56+v27<<(uint(int32(2))%32)))) = v60
				v63 = v27 << (uint(int32(1)) % 32)
				v65 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayRemove[3]))
				v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
				*(*uint8)(unsafe.Add(mBase, uint32(v63+v66)+1)) = uint8(v60)
				v71 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayRemove[3]))
				v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+8))
				*(*uint8)(unsafe.Add(mBase, uint32(v72+v63))) = uint8(v60)
			} else {
			}
			v81 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayRemove[3]))
			v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+12))
			v84 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v82+v27))) = uint8(v84)
			v86 = int32(2)
			v87 = v27 << (uint(v86) % 32)
			v89 = v12 + int32(36)
			v91 = v27 + int32(1)
			v93 = v91 << (uint(v86) % 32)
			v94 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
			v97 = v94 + (v27 ^ int32(-1))
			v99 = v97 << (uint(v86) % 32)
			v101 = base.B2i32(v99 == v84)
			if v101 == v84 {
				base.MemoryCopy(m, v89+v87, v89+v93, v99)
			} else {
			}
			if v101 == int32(0) {
				v110 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayRemove[3]))
				v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
				base.MemoryCopy(m, v87+v111, v111+v93, v99)
			} else {
			}
			v117 = v97 << (uint(int32(1)) % 32)
			if v117 != 0 {
				v119 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayRemove[3]))
				v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+8))
				v121 = int32(1)
				base.MemoryCopy(m, v120+v27<<(uint(v121)%32), v120+v91<<(uint(v121)%32), v117)
			} else {
			}
			if v97 != 0 {
				v130 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayRemove[3]))
				v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)+12))
				base.MemoryCopy(m, v131+v27, v131+v91, v97)
			} else {
			}
			v136 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
			*(*int32)(unsafe.Add(mBase, uint32(v89+v136<<(uint(int32(2))%32)-int32(4)))) = int32(-1)
			v144 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
			v146 = v144 - int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v12))) = v146
			if v27 < v146 {
				v150 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayRemove[4]))
				v151 = v27
				for {
					v164 = *(*int32)(unsafe.Add(mBase, uint32(v89+v151<<(uint(int32(2))%32))))
					*(*int32)(unsafe.Add(mBase, uint32(v150+v164*int32(640))+48)) = v151
					v170 = v151 + int32(1)
					v171 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					if v170 < v171 {
						v151 = v170
						continue
					} else {
						break
					}
					break
				}
			} else {
			}
			v184 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayRemove[1]))
			F_LWLockRelease(m, v184+int32(384))
			mBase = m.M
			v188 = m.ExcPending
			if v188 != 0 {
				return
			} else {
				v190 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayRemove[1]))
				F_LWLockRelease(m, v190+int32(512))
				mBase = m.M
				v194 = m.ExcPending
				if v194 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_ProcNumberGetProc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	v2 = int32(0)
	if l0 < v2 {
		v20 = v2
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, _c_F_ProcNumberGetProc[0]))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
		if base.Ui32(v9) <= base.Ui32(l0) {
			v20 = int32(0)
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
			v14 = v11 + l0*int32(640)
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
			if v16 != 0 {
				v17 = v14
			} else {
				v17 = int32(0)
			}
			v20 = v17
		}
	}
	return v20
}
func F_proc_exit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_proc_exit[0]))
	if v8 == int32(42) {
		F_proc_exit_prepare(m, l0)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			v15 = F_errstart(m, int32(12), int32(0))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				if v15 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(v5))) = l0
					F_errmsg_internal(m, int32(_a_F_proc_exit_0), v5)
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_proc_exit_1), int32(155), int32(_a_F_proc_exit_2))
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return
						} else {
							F_pgl_exit(m, l0)
							mBase = m.M
							v27 = m.ExcPending
							if v27 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					F_pgl_exit(m, l0)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(23), int32(0))
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(_a_F_proc_exit_3), int32(0))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_proc_exit_1), int32(109), int32(_a_F_proc_exit_2))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
