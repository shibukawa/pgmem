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
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int64
	_ = v32
	var v33 int32
	_ = v33
	var v47 int64
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
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
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			if l1 != 0 {
				v28 = int32(3)
				v31 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayRemove[2]))
				v32 = *(*int64)(unsafe.Add(mBase, uint32(v31)+48))
				v33 = base.I32_wrap_i64(v32)
				if base.B2i32(base.Ui32(l1) < base.Ui32(v28))|base.B2i32(base.Ui32(v33) < base.Ui32(v28)) == int32(0) {
					if v33-l1 < int32(0) {
						*(*int64)(unsafe.Add(mBase, uint32(v31)+48)) = v32 + base.I64_extend_i32_s(l1-v33)
					} else {
					}
				} else {
					if base.Ui32(l1) <= base.Ui32(v33) {
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v31)+48)) = v32 + base.I64_extend_i32_s(l1-v33)
					}
				}
				v47 = *(*int64)(unsafe.Add(mBase, uint32(v31)+56))
				*(*int64)(unsafe.Add(mBase, uint32(v31)+56)) = v47 + int64(1)
				v51 = int32(_a_F_ProcArrayRemove_0)
				v52 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayRemove[3]))
				v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
				v57 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v53+v27<<(uint(int32(2))%32)))) = v57
				v60 = v27 << (uint(int32(1)) % 32)
				v62 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayRemove[3]))
				v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
				*(*uint8)(unsafe.Add(mBase, uint32(v60+v63)+1)) = uint8(v57)
				v68 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayRemove[3]))
				v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+8))
				*(*uint8)(unsafe.Add(mBase, uint32(v69+v60))) = uint8(v57)
			} else {
			}
			v78 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayRemove[3]))
			v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
			v81 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v79+v27))) = uint8(v81)
			v83 = int32(2)
			v84 = v27 << (uint(v83) % 32)
			v86 = v12 + int32(36)
			v88 = v27 + int32(1)
			v90 = v88 << (uint(v83) % 32)
			v91 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
			v94 = v91 + (v27 ^ int32(-1))
			v96 = v94 << (uint(v83) % 32)
			v98 = base.B2i32(v96 == v81)
			if v98 == v81 {
				base.MemoryCopy(m, v86+v84, v86+v90, v96)
			} else {
			}
			if v98 == int32(0) {
				v107 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayRemove[3]))
				v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
				base.MemoryCopy(m, v84+v108, v108+v90, v96)
			} else {
			}
			v114 = v94 << (uint(int32(1)) % 32)
			if v114 != 0 {
				v116 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayRemove[3]))
				v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+8))
				v118 = int32(1)
				base.MemoryCopy(m, v117+v27<<(uint(v118)%32), v117+v88<<(uint(v118)%32), v114)
			} else {
			}
			if v94 != 0 {
				v127 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayRemove[3]))
				v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)+12))
				base.MemoryCopy(m, v128+v27, v128+v88, v94)
			} else {
			}
			v133 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
			*(*int32)(unsafe.Add(mBase, uint32(v86+v133<<(uint(int32(2))%32)-int32(4)))) = int32(-1)
			v141 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
			v143 = v141 - int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v12))) = v143
			if v27 < v143 {
				v147 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayRemove[4]))
				v148 = v27
				for {
					v161 = *(*int32)(unsafe.Add(mBase, uint32(v86+v148<<(uint(int32(2))%32))))
					*(*int32)(unsafe.Add(mBase, uint32(v147+v161*int32(768))+32)) = v148
					v167 = v148 + int32(1)
					v168 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					if v167 < v168 {
						v148 = v167
						continue
					} else {
						break
					}
					break
				}
			} else {
			}
			v181 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayRemove[1]))
			F_LWLockRelease(m, v181+int32(384))
			mBase = m.M
			v185 = m.ExcPending
			if v185 != 0 {
				return
			} else {
				v187 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayRemove[1]))
				F_LWLockRelease(m, v187+int32(512))
				mBase = m.M
				v191 = m.ExcPending
				if v191 != 0 {
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
			v14 = v11 + l0*int32(768)
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
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
	var v9 int32
	_ = v9
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
	v9 = m.Env.Pgmem_getpid(m)
	mBase = m.M
	if v8 == v9 {
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
		F_errstart_cold(m, int32(24), int32(0))
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
