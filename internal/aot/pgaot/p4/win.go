package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_WinRowsArePeers(m *base.Module, l0 int32, l1 int64, l2 int64) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+96))
	if v15 == int32(0) {
		v72 = int32(1)
		m.G0 = v11 + int32(32)
		return v72
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v13)+404))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)+400))
		v21 = F_window_gettupleslot(m, l0, l1, v20)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			if v21 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v83 = m.ExcPending
				if v83 != 0 {
					return int32(0)
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = l1
					F_errmsg_internal(m, int32(_a_F_WinRowsArePeers_0), v11+int32(16))
					mBase = m.M
					v89 = m.ExcPending
					if v89 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_WinRowsArePeers_1), int32(3344), int32(_a_F_WinRowsArePeers_2))
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v27 = F_window_gettupleslot(m, l0, l2, v19)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					if v27 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v98 = m.ExcPending
						if v98 != 0 {
							return int32(0)
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v11))) = l2
							F_errmsg_internal(m, int32(_a_F_WinRowsArePeers_0), v11)
							mBase = m.M
							v102 = m.ExcPending
							if v102 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_WinRowsArePeers_1), int32(3347), int32(_a_F_WinRowsArePeers_2))
								mBase = m.M
								v107 = m.ExcPending
								if v107 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+96))
						if v32 == int32(0) {
							v61 = int32(1)
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
							v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
							m.T0[v64].(func(*base.Module, int32))(m, v20)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								v67 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
								v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
								m.T0[v68].(func(*base.Module, int32))(m, v19)
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return int32(0)
								} else {
									v72 = v61
									m.G0 = v11 + int32(32)
									return v72
								}
							}
						} else {
							v36 = *(*int32)(unsafe.Add(mBase, uint32(v13)+372))
							*(*int32)(unsafe.Add(mBase, uint32(v36)+8)) = v19
							*(*int32)(unsafe.Add(mBase, uint32(v36)+12)) = v20
							v39 = *(*int32)(unsafe.Add(mBase, uint32(v13)+140))
							if v39 != 0 {
								v40 = int32(_a_F_WinRowsArePeers_3)
								v41 = *(*int32)(unsafe.Add(mBase, _c_F_WinRowsArePeers[0]))
								v43 = *(*int32)(unsafe.Add(mBase, uint32(v36)+20))
								*(*int32)(unsafe.Add(mBase, _c_F_WinRowsArePeers[0])) = v43
								v47 = *(*int32)(unsafe.Add(mBase, uint32(v39)+20))
								v48 = m.T0[v47].(func(*base.Module, int32, int32, int32) int32)(m, v39, v36, v11+int32(31))
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_WinRowsArePeers[0])) = v41
									v56 = base.B2i32(v48 != int32(0))
									v57 = *(*int32)(unsafe.Add(mBase, uint32(v36)+20))
									F_MemoryContextReset(m, v57)
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
										return int32(0)
									} else {
										v61 = v56
										v63 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
										v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
										m.T0[v64].(func(*base.Module, int32))(m, v20)
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return int32(0)
										} else {
											v67 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
											v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
											m.T0[v68].(func(*base.Module, int32))(m, v19)
											mBase = m.M
											v70 = m.ExcPending
											if v70 != 0 {
												return int32(0)
											} else {
												v72 = v61
												m.G0 = v11 + int32(32)
												return v72
											}
										}
									}
								}
							} else {
								v56 = int32(1)
								v57 = *(*int32)(unsafe.Add(mBase, uint32(v36)+20))
								F_MemoryContextReset(m, v57)
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return int32(0)
								} else {
									v61 = v56
									v63 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
									v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
									m.T0[v64].(func(*base.Module, int32))(m, v20)
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return int32(0)
									} else {
										v67 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
										v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
										m.T0[v68].(func(*base.Module, int32))(m, v19)
										mBase = m.M
										v70 = m.ExcPending
										if v70 != 0 {
											return int32(0)
										} else {
											v72 = v61
											m.G0 = v11 + int32(32)
											return v72
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_win_to_utf8(m *base.Module, l0 int32) int32 {
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v10 = Fn13928(m, l0, int32(_a_F_win_to_utf8_0), int32(_a_F_win_to_utf8_1), int32(114), int32(_a_F_win_to_utf8_2), int32(_a_F_win_to_utf8_3), int32(_a_F_win_to_utf8_4), int32(15), int32(18))
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		return v10
	}
}
