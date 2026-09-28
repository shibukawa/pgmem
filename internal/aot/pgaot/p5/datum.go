package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_DatumGetEOHP(m *base.Module, l0 int64) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(l0))+2))
	return v3
}
func F_DatumGetExpandedArray(m *base.Module, l0 int64) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v19 int32
	_ = v19
	v3 = base.I32_wrap_i64(l0)
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3))))
	if v4 == int32(1) {
		v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+1)))
		if v7 == int32(3) {
			v17 = l0
			v19 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v17))+2))
			return v19
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, _c_F_DatumGetExpandedArray[0]))
			v13 = F_expand_array(m, l0, v11, int32(0))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				v17 = v13
				v19 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v17))+2))
				return v19
			}
		}
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, _c_F_DatumGetExpandedArray[0]))
		v13 = F_expand_array(m, l0, v11, int32(0))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = v13
			v19 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v17))+2))
			return v19
		}
	}
}
func F_datum_write(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v199 int32
	_ = v199
	v2 = l1
	v8 = m.G0
	v10 = v8 + int32(-64)
	m.G0 = v10
	if l2 != 0 {
		switch l3 - int32(99) {
		case 0:
			v34 = l0
			v35 = int32(-1)
			if base.I32_popcnt(l4) != int32(1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l4
					F_errmsg_internal(m, int32(_a_F_datum_write_0), v8+int32(-48))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_datum_write_1), int32(474), int32(_a_F_datum_write_2))
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v39 = v34 & v35
				switch base.I32_ctz(l4) {
				case 0:
					*(*uint8)(unsafe.Add(mBase, uint32(v39))) = uint8(v2)
					v169 = v39
					v172 = l4
					m.G0 = v10 - int32(-64)
					return v169 + v172
				case 1:
					*(*uint16)(unsafe.Add(mBase, uint32(v39))) = uint16(v2)
					v169 = v39
					v172 = l4
					m.G0 = v10 - int32(-64)
					return v169 + v172
				case 2:
					*(*uint32)(unsafe.Add(mBase, uint32(v39))) = uint32(v2)
					v169 = v39
					v172 = l4
					m.G0 = v10 - int32(-64)
					return v169 + v172
				case 3:
					*(*int64)(unsafe.Add(mBase, uint32(v39))) = v2
					v169 = v39
					v172 = l4
					m.G0 = v10 - int32(-64)
					return v169 + v172
				default:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l4
						F_errmsg_internal(m, int32(_a_F_datum_write_0), v8+int32(-48))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_datum_write_1), int32(474), int32(_a_F_datum_write_2))
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
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
		case 1:
			v34 = l0 + int32(7)
			v35 = int32(-8)
			if base.I32_popcnt(l4) != int32(1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l4
					F_errmsg_internal(m, int32(_a_F_datum_write_0), v8+int32(-48))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_datum_write_1), int32(474), int32(_a_F_datum_write_2))
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v39 = v34 & v35
				switch base.I32_ctz(l4) {
				case 0:
					*(*uint8)(unsafe.Add(mBase, uint32(v39))) = uint8(v2)
					v169 = v39
					v172 = l4
					m.G0 = v10 - int32(-64)
					return v169 + v172
				case 1:
					*(*uint16)(unsafe.Add(mBase, uint32(v39))) = uint16(v2)
					v169 = v39
					v172 = l4
					m.G0 = v10 - int32(-64)
					return v169 + v172
				case 2:
					*(*uint32)(unsafe.Add(mBase, uint32(v39))) = uint32(v2)
					v169 = v39
					v172 = l4
					m.G0 = v10 - int32(-64)
					return v169 + v172
				case 3:
					*(*int64)(unsafe.Add(mBase, uint32(v39))) = v2
					v169 = v39
					v172 = l4
					m.G0 = v10 - int32(-64)
					return v169 + v172
				default:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l4
						F_errmsg_internal(m, int32(_a_F_datum_write_0), v8+int32(-48))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_datum_write_1), int32(474), int32(_a_F_datum_write_2))
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
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
		default:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = l3
				F_errmsg_internal(m, int32(_a_F_datum_write_3), v10)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_datum_write_1), int32(322), int32(_a_F_datum_write_4))
					mBase = m.M
					v199 = m.ExcPending
					if v199 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		case 6:
			v34 = l0 + int32(3)
			v35 = int32(-4)
			if base.I32_popcnt(l4) != int32(1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l4
					F_errmsg_internal(m, int32(_a_F_datum_write_0), v8+int32(-48))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_datum_write_1), int32(474), int32(_a_F_datum_write_2))
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v39 = v34 & v35
				switch base.I32_ctz(l4) {
				case 0:
					*(*uint8)(unsafe.Add(mBase, uint32(v39))) = uint8(v2)
					v169 = v39
					v172 = l4
					m.G0 = v10 - int32(-64)
					return v169 + v172
				case 1:
					*(*uint16)(unsafe.Add(mBase, uint32(v39))) = uint16(v2)
					v169 = v39
					v172 = l4
					m.G0 = v10 - int32(-64)
					return v169 + v172
				case 2:
					*(*uint32)(unsafe.Add(mBase, uint32(v39))) = uint32(v2)
					v169 = v39
					v172 = l4
					m.G0 = v10 - int32(-64)
					return v169 + v172
				case 3:
					*(*int64)(unsafe.Add(mBase, uint32(v39))) = v2
					v169 = v39
					v172 = l4
					m.G0 = v10 - int32(-64)
					return v169 + v172
				default:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l4
						F_errmsg_internal(m, int32(_a_F_datum_write_0), v8+int32(-48))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_datum_write_1), int32(474), int32(_a_F_datum_write_2))
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
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
		case 16:
			v34 = l0 + int32(1)
			v35 = int32(-2)
			if base.I32_popcnt(l4) != int32(1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l4
					F_errmsg_internal(m, int32(_a_F_datum_write_0), v8+int32(-48))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_datum_write_1), int32(474), int32(_a_F_datum_write_2))
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v39 = v34 & v35
				switch base.I32_ctz(l4) {
				case 0:
					*(*uint8)(unsafe.Add(mBase, uint32(v39))) = uint8(v2)
					v169 = v39
					v172 = l4
					m.G0 = v10 - int32(-64)
					return v169 + v172
				case 1:
					*(*uint16)(unsafe.Add(mBase, uint32(v39))) = uint16(v2)
					v169 = v39
					v172 = l4
					m.G0 = v10 - int32(-64)
					return v169 + v172
				case 2:
					*(*uint32)(unsafe.Add(mBase, uint32(v39))) = uint32(v2)
					v169 = v39
					v172 = l4
					m.G0 = v10 - int32(-64)
					return v169 + v172
				case 3:
					*(*int64)(unsafe.Add(mBase, uint32(v39))) = v2
					v169 = v39
					v172 = l4
					m.G0 = v10 - int32(-64)
					return v169 + v172
				default:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l4
						F_errmsg_internal(m, int32(_a_F_datum_write_0), v8+int32(-48))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_datum_write_1), int32(474), int32(_a_F_datum_write_2))
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
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
		}
	} else {
		switch l4 + int32(2) {
		case 0:
			v133 = base.I32_wrap_i64(v2)
			v134 = F_strlen(m, v133)
			mBase = m.M
			v136 = v134 + int32(1)
			if v136 == int32(0) {
				v169 = l0
				v172 = v136
			} else {
				base.MemoryCopy(m, l0, v133, v136)
				v169 = l0
				v172 = v136
			}
			m.G0 = v10 - int32(-64)
			return v169 + v172
		case 1:
			v63 = base.I32_wrap_i64(v2)
			v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
			if v64 == int32(1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v182 = m.ExcPending
				if v182 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_datum_write_5), int32(0))
					mBase = m.M
					v186 = m.ExcPending
					if v186 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_datum_write_6), int32(2966), int32(_a_F_datum_write_7))
						mBase = m.M
						v191 = m.ExcPending
						if v191 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				if v64&int32(1) != 0 {
					v70 = int32(base.Ui32(v64) >> (uint(int32(1)) % 32))
					if v70 == int32(0) {
						v169 = l0
						v172 = v70
					} else {
						base.MemoryCopy(m, l0, v63, v70)
						v169 = l0
						v172 = v70
					}
					m.G0 = v10 - int32(-64)
					return v169 + v172
				} else {
					if v64&int32(2)|base.B2i32(l5 == int32(112)) != 0 {
						switch l3 - int32(99) {
						case 0:
							v124 = l0
							v125 = int32(-1)
							v126 = v124 & v125
							v127 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
							v129 = int32(base.Ui32(v127) >> (uint(int32(2)) % 32))
							if v129 == int32(0) {
								v169 = v126
								v172 = v129
							} else {
								base.MemoryCopy(m, v126, v63, v129)
								v169 = v126
								v172 = v129
							}
							m.G0 = v10 - int32(-64)
							return v169 + v172
						case 1:
							v124 = l0 + int32(7)
							v125 = int32(-8)
							v126 = v124 & v125
							v127 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
							v129 = int32(base.Ui32(v127) >> (uint(int32(2)) % 32))
							if v129 == int32(0) {
								v169 = v126
								v172 = v129
							} else {
								base.MemoryCopy(m, v126, v63, v129)
								v169 = v126
								v172 = v129
							}
							m.G0 = v10 - int32(-64)
							return v169 + v172
						default:
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v108 = m.ExcPending
							if v108 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = l3
								F_errmsg_internal(m, int32(_a_F_datum_write_3), v8+int32(-16))
								mBase = m.M
								v114 = m.ExcPending
								if v114 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_datum_write_1), int32(322), int32(_a_F_datum_write_4))
									mBase = m.M
									v199 = m.ExcPending
									if v199 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						case 6:
							v124 = l0 + int32(3)
							v125 = int32(-4)
							v126 = v124 & v125
							v127 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
							v129 = int32(base.Ui32(v127) >> (uint(int32(2)) % 32))
							if v129 == int32(0) {
								v169 = v126
								v172 = v129
							} else {
								base.MemoryCopy(m, v126, v63, v129)
								v169 = v126
								v172 = v129
							}
							m.G0 = v10 - int32(-64)
							return v169 + v172
						case 16:
							v124 = l0 + int32(1)
							v125 = int32(-2)
							v126 = v124 & v125
							v127 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
							v129 = int32(base.Ui32(v127) >> (uint(int32(2)) % 32))
							if v129 == int32(0) {
								v169 = v126
								v172 = v129
							} else {
								base.MemoryCopy(m, v126, v63, v129)
								v169 = v126
								v172 = v129
							}
							m.G0 = v10 - int32(-64)
							return v169 + v172
						}
					} else {
						v79 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
						v81 = int32(base.Ui32(v79) >> (uint(int32(2)) % 32))
						v83 = v81 - int32(3)
						if base.Ui32(int32(127)) < base.Ui32(v83) {
							switch l3 - int32(99) {
							case 0:
								v124 = l0
								v125 = int32(-1)
								v126 = v124 & v125
								v127 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
								v129 = int32(base.Ui32(v127) >> (uint(int32(2)) % 32))
								if v129 == int32(0) {
									v169 = v126
									v172 = v129
								} else {
									base.MemoryCopy(m, v126, v63, v129)
									v169 = v126
									v172 = v129
								}
								m.G0 = v10 - int32(-64)
								return v169 + v172
							case 1:
								v124 = l0 + int32(7)
								v125 = int32(-8)
								v126 = v124 & v125
								v127 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
								v129 = int32(base.Ui32(v127) >> (uint(int32(2)) % 32))
								if v129 == int32(0) {
									v169 = v126
									v172 = v129
								} else {
									base.MemoryCopy(m, v126, v63, v129)
									v169 = v126
									v172 = v129
								}
								m.G0 = v10 - int32(-64)
								return v169 + v172
							default:
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v108 = m.ExcPending
								if v108 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = l3
									F_errmsg_internal(m, int32(_a_F_datum_write_3), v8+int32(-16))
									mBase = m.M
									v114 = m.ExcPending
									if v114 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_datum_write_1), int32(322), int32(_a_F_datum_write_4))
										mBase = m.M
										v199 = m.ExcPending
										if v199 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							case 6:
								v124 = l0 + int32(3)
								v125 = int32(-4)
								v126 = v124 & v125
								v127 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
								v129 = int32(base.Ui32(v127) >> (uint(int32(2)) % 32))
								if v129 == int32(0) {
									v169 = v126
									v172 = v129
								} else {
									base.MemoryCopy(m, v126, v63, v129)
									v169 = v126
									v172 = v129
								}
								m.G0 = v10 - int32(-64)
								return v169 + v172
							case 16:
								v124 = l0 + int32(1)
								v125 = int32(-2)
								v126 = v124 & v125
								v127 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
								v129 = int32(base.Ui32(v127) >> (uint(int32(2)) % 32))
								if v129 == int32(0) {
									v169 = v126
									v172 = v129
								} else {
									base.MemoryCopy(m, v126, v63, v129)
									v169 = v126
									v172 = v129
								}
								m.G0 = v10 - int32(-64)
								return v169 + v172
							}
						} else {
							v86 = int32(1)
							v89 = v83<<(uint(v86)%32) | v86
							*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v89)
							v92 = v81 - int32(4)
							if v92 == int32(0) {
								v169 = l0
								v172 = v83
							} else {
								base.MemoryCopy(m, l0+int32(1), v63+int32(4), v92)
								v169 = l0
								v172 = v83
							}
							m.G0 = v10 - int32(-64)
							return v169 + v172
						}
					}
				}
			}
		default:
			switch l3 - int32(99) {
			case 0:
				v162 = l0
				v163 = int32(-1)
				v164 = v162 & v163
				if l4 == int32(0) {
					v169 = v164
					v172 = l4
				} else {
					base.MemoryCopy(m, v164, base.I32_wrap_i64(v2), l4)
					v169 = v164
					v172 = l4
				}
				m.G0 = v10 - int32(-64)
				return v169 + v172
			case 1:
				v162 = l0 + int32(7)
				v163 = int32(-8)
				v164 = v162 & v163
				if l4 == int32(0) {
					v169 = v164
					v172 = l4
				} else {
					base.MemoryCopy(m, v164, base.I32_wrap_i64(v2), l4)
					v169 = v164
					v172 = l4
				}
				m.G0 = v10 - int32(-64)
				return v169 + v172
			default:
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v146 = m.ExcPending
				if v146 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = l3
					F_errmsg_internal(m, int32(_a_F_datum_write_3), v8+int32(-32))
					mBase = m.M
					v152 = m.ExcPending
					if v152 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_datum_write_1), int32(322), int32(_a_F_datum_write_4))
						mBase = m.M
						v199 = m.ExcPending
						if v199 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			case 6:
				v162 = l0 + int32(3)
				v163 = int32(-4)
				v164 = v162 & v163
				if l4 == int32(0) {
					v169 = v164
					v172 = l4
				} else {
					base.MemoryCopy(m, v164, base.I32_wrap_i64(v2), l4)
					v169 = v164
					v172 = l4
				}
				m.G0 = v10 - int32(-64)
				return v169 + v172
			case 16:
				v162 = l0 + int32(1)
				v163 = int32(-2)
				v164 = v162 & v163
				if l4 == int32(0) {
					v169 = v164
					v172 = l4
				} else {
					base.MemoryCopy(m, v164, base.I32_wrap_i64(v2), l4)
					v169 = v164
					v172 = l4
				}
				m.G0 = v10 - int32(-64)
				return v169 + v172
			}
		}
	}
}
