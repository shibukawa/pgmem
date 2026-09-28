package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_dependencies_object_start(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	v5 = m.G0
	v7 = v5 - int32(160)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v9 {
	case 0:
		v10 = int32(23)
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v12 = F_errsave_start(m, v11)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			if v12 == int32(0) {
				v190 = v10
				m.G0 = v7 + int32(160)
				return v190
			} else {
				F_errcode(m, int32(33685634))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v21
					F_errmsg(m, int32(_a_F_dependencies_object_start_0), v7+int32(16))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						v30 = F_errdetail(m, int32(_a_F_dependencies_object_start_1), int32(0))
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							F_errsave_finish(m, v11, int32(_a_F_dependencies_object_start_2), int32(80), int32(_a_F_dependencies_object_start_3))
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								v190 = v10
								m.G0 = v7 + int32(160)
								return v190
							}
						}
					}
				}
			}
		}
	case 1:
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(2)
		v190 = int32(0)
		m.G0 = v7 + int32(160)
		return v190
	case 2:
		v37 = int32(23)
		v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v39 = F_errsave_start(m, v38)
		mBase = m.M
		v40 = m.ExcPending
		if v40 != 0 {
			return int32(0)
		} else {
			if v39 == int32(0) {
				v190 = v37
				m.G0 = v7 + int32(160)
				return v190
			} else {
				F_errcode(m, int32(33685634))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v46
					F_errmsg(m, int32(_a_F_dependencies_object_start_0), v7+int32(32))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						v55 = F_errdetail(m, int32(_a_F_dependencies_object_start_4), int32(0))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							F_errsave_finish(m, v38, int32(_a_F_dependencies_object_start_2), int32(88), int32(_a_F_dependencies_object_start_3))
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								v190 = v37
								m.G0 = v7 + int32(160)
								return v190
							}
						}
					}
				}
			}
		}
	case 3:
		v62 = int32(23)
		v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v64 = F_errsave_start(m, v63)
		mBase = m.M
		v65 = m.ExcPending
		if v65 != 0 {
			return int32(0)
		} else {
			if v64 == int32(0) {
				v190 = v62
				m.G0 = v7 + int32(160)
				return v190
			} else {
				F_errcode(m, int32(33685634))
				mBase = m.M
				v70 = m.ExcPending
				if v70 != 0 {
					return int32(0)
				} else {
					v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v7)+64)) = v71
					F_errmsg(m, int32(_a_F_dependencies_object_start_0), v7-int32(-64))
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = int32(_a_F_dependencies_object_start_5)
						v83 = F_errdetail(m, int32(_a_F_dependencies_object_start_6), v7+int32(48))
						mBase = m.M
						v84 = m.ExcPending
						if v84 != 0 {
							return int32(0)
						} else {
							F_errsave_finish(m, v63, int32(_a_F_dependencies_object_start_2), int32(97), int32(_a_F_dependencies_object_start_3))
							mBase = m.M
							v89 = m.ExcPending
							if v89 != 0 {
								return int32(0)
							} else {
								v190 = v62
								m.G0 = v7 + int32(160)
								return v190
							}
						}
					}
				}
			}
		}
	case 4:
		v90 = int32(23)
		v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v92 = F_errsave_start(m, v91)
		mBase = m.M
		v93 = m.ExcPending
		if v93 != 0 {
			return int32(0)
		} else {
			if v92 == int32(0) {
				v190 = v90
				m.G0 = v7 + int32(160)
				return v190
			} else {
				F_errcode(m, int32(33685634))
				mBase = m.M
				v98 = m.ExcPending
				if v98 != 0 {
					return int32(0)
				} else {
					v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v7)+80)) = v99
					F_errmsg(m, int32(_a_F_dependencies_object_start_0), v7+int32(80))
					mBase = m.M
					v105 = m.ExcPending
					if v105 != 0 {
						return int32(0)
					} else {
						v108 = F_errdetail(m, int32(_a_F_dependencies_object_start_7), int32(0))
						mBase = m.M
						v109 = m.ExcPending
						if v109 != 0 {
							return int32(0)
						} else {
							F_errsave_finish(m, v91, int32(_a_F_dependencies_object_start_2), int32(105), int32(_a_F_dependencies_object_start_3))
							mBase = m.M
							v114 = m.ExcPending
							if v114 != 0 {
								return int32(0)
							} else {
								v190 = v90
								m.G0 = v7 + int32(160)
								return v190
							}
						}
					}
				}
			}
		}
	case 5:
		v115 = int32(23)
		v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v117 = F_errsave_start(m, v116)
		mBase = m.M
		v118 = m.ExcPending
		if v118 != 0 {
			return int32(0)
		} else {
			if v117 == int32(0) {
				v190 = v115
				m.G0 = v7 + int32(160)
				return v190
			} else {
				F_errcode(m, int32(33685634))
				mBase = m.M
				v123 = m.ExcPending
				if v123 != 0 {
					return int32(0)
				} else {
					v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v7)+112)) = v124
					F_errmsg(m, int32(_a_F_dependencies_object_start_0), v7+int32(112))
					mBase = m.M
					v130 = m.ExcPending
					if v130 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+96)) = int32(_a_F_dependencies_object_start_8)
						v136 = F_errdetail(m, int32(_a_F_dependencies_object_start_9), v7+int32(96))
						mBase = m.M
						v137 = m.ExcPending
						if v137 != 0 {
							return int32(0)
						} else {
							F_errsave_finish(m, v116, int32(_a_F_dependencies_object_start_2), int32(114), int32(_a_F_dependencies_object_start_3))
							mBase = m.M
							v142 = m.ExcPending
							if v142 != 0 {
								return int32(0)
							} else {
								v190 = v115
								m.G0 = v7 + int32(160)
								return v190
							}
						}
					}
				}
			}
		}
	case 6:
		v143 = int32(23)
		v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v145 = F_errsave_start(m, v144)
		mBase = m.M
		v146 = m.ExcPending
		if v146 != 0 {
			return int32(0)
		} else {
			if v145 == int32(0) {
				v190 = v143
				m.G0 = v7 + int32(160)
				return v190
			} else {
				F_errcode(m, int32(33685634))
				mBase = m.M
				v151 = m.ExcPending
				if v151 != 0 {
					return int32(0)
				} else {
					v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v7)+144)) = v152
					F_errmsg(m, int32(_a_F_dependencies_object_start_0), v7+int32(144))
					mBase = m.M
					v158 = m.ExcPending
					if v158 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+128)) = int32(_a_F_dependencies_object_start_10)
						v164 = F_errdetail(m, int32(_a_F_dependencies_object_start_9), v7+int32(128))
						mBase = m.M
						v165 = m.ExcPending
						if v165 != 0 {
							return int32(0)
						} else {
							F_errsave_finish(m, v144, int32(_a_F_dependencies_object_start_2), int32(123), int32(_a_F_dependencies_object_start_3))
							mBase = m.M
							v170 = m.ExcPending
							if v170 != 0 {
								return int32(0)
							} else {
								v190 = v143
								m.G0 = v7 + int32(160)
								return v190
							}
						}
					}
				}
			}
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v174 = m.ExcPending
		if v174 != 0 {
			return int32(0)
		} else {
			v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v175
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_dependencies_object_start_11)
			F_errmsg_internal(m, int32(_a_F_dependencies_object_start_12), v7)
			mBase = m.M
			v181 = m.ExcPending
			if v181 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_dependencies_object_start_2), int32(129), int32(_a_F_dependencies_object_start_3))
				mBase = m.M
				v186 = m.ExcPending
				if v186 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
