package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_KnownAssignedTransactionIdsIdleMaintenance(m *base.Module) {
	var v4 int32
	_ = v4
	F_KnownAssignedXidsCompress(m, int32(3), int32(0))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_KnownAssignedXidsAdd(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v198 int32
	_ = v198
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsAdd[0]))
	if base.Ui32(l1) < base.Ui32(l0) {
		v15 = int32(1)
		if base.B2i32(base.Ui32(int32(2)) < base.Ui32(l1))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(l0)) == int32(0) {
			v27 = base.B2i32(base.Ui32(l0) < base.Ui32(l1))
		} else {
			v27 = int32(base.Ui32(l0-l1) >> (uint(int32(31)) % 32))
		}
		if v27 == int32(0) {
			v68 = v15
		} else {
			v33 = l0
			v34 = v15
			for {
				v41 = int32(1)
				v42 = v34 + v41
				v43 = int32(3)
				v45 = v33 + v41
				if base.Ui32(v45) <= base.Ui32(v43) {
					v48 = v43
				} else {
					v48 = v45
				}
				if base.B2i32(base.Ui32(int32(2)) < base.Ui32(l1))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v48)) == int32(0) {
					v60 = base.B2i32(base.Ui32(v48) < base.Ui32(l1))
				} else {
					v60 = int32(base.Ui32(v48-l1) >> (uint(int32(31)) % 32))
				}
				if v60 != 0 {
					v33 = v48
					v34 = v42
					continue
				} else {
					break
				}
				break
			}
			v68 = v42
		}
	} else {
		v68 = l1 - l0 + int32(1)
	}
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	if v76 < v75 {
		v79 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsAdd[1]))
		v80 = int32(2)
		v85 = *(*int32)(unsafe.Add(mBase, uint32(v79+v75<<(uint(v80)%32)-int32(4))))
		if base.B2i32(base.Ui32(v80) < base.Ui32(l0))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v85)) == int32(0) {
			v97 = base.B2i32(base.Ui32(l0) <= base.Ui32(v85))
		} else {
			v97 = base.B2i32(int32(0) <= v85-l0)
		}
		if v97 != 0 {
			F_KnownAssignedXidsDisplay(m, int32(15))
			mBase = m.M
			v204 = m.ExcPending
			if v204 != 0 {
				return
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v208 = m.ExcPending
				if v208 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(_a_F_KnownAssignedXidsAdd_0), int32(0))
					mBase = m.M
					v212 = m.ExcPending
					if v212 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_KnownAssignedXidsAdd_1), int32(_a_F_KnownAssignedXidsAdd_2), int32(_a_F_KnownAssignedXidsAdd_3))
						mBase = m.M
						v217 = m.ExcPending
						if v217 != 0 {
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
			v98 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
			if v98 < v75+v68 {
				F_KnownAssignedXidsCompress(m, int32(0), l2)
				mBase = m.M
				v103 = m.ExcPending
				if v103 != 0 {
					return
				} else {
					v104 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
					v105 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
					if v104 < v105+v68 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v221 = m.ExcPending
						if v221 != 0 {
							return
						} else {
							F_errmsg_internal(m, int32(_a_F_KnownAssignedXidsAdd_4), int32(0))
							mBase = m.M
							v225 = m.ExcPending
							if v225 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_KnownAssignedXidsAdd_1), int32(_a_F_KnownAssignedXidsAdd_5), int32(_a_F_KnownAssignedXidsAdd_3))
								mBase = m.M
								v230 = m.ExcPending
								if v230 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v108 = v105
						if v68 <= int32(0) {
							v190 = v108
						} else {
							v112 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsAdd[2]))
							v114 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsAdd[1]))
							if v68 != int32(1) {
								v121 = l0
								v124 = v108
								v131 = int32(0)
								for {
									v132 = int32(2)
									*(*int32)(unsafe.Add(mBase, uint32(v114+v124<<(uint(v132)%32)))) = v121
									v137 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v112+v124))) = uint8(v137)
									v140 = v124 + v137
									v144 = int32(3)
									v146 = v121 + v137
									if base.Ui32(v146) <= base.Ui32(v144) {
										v149 = v144
									} else {
										v149 = v146
									}
									*(*int32)(unsafe.Add(mBase, uint32(v114+v140<<(uint(v132)%32)))) = v149
									v152 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v112+v140))) = uint8(v152)
									v154 = int32(3)
									v156 = v149 + v152
									if base.Ui32(v156) <= base.Ui32(v154) {
										v159 = v154
									} else {
										v159 = v156
									}
									v160 = int32(2)
									v161 = v124 + v160
									v163 = v131 + v160
									if v163 != v68&int32(2147483646) {
										v121 = v159
										v124 = v161
										v131 = v163
										continue
									} else {
										break
									}
									break
								}
								if v68&int32(1) == int32(0) {
									v190 = v161
								} else {
									v167 = v159
									v170 = v161
									*(*int32)(unsafe.Add(mBase, uint32(v114+v170<<(uint(int32(2))%32)))) = v167
									v183 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v112+v170))) = uint8(v183)
									v190 = v170 + v183
								}
							} else {
								v167 = l0
								v170 = v108
								*(*int32)(unsafe.Add(mBase, uint32(v114+v170<<(uint(int32(2))%32)))) = v167
								v183 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v112+v170))) = uint8(v183)
								v190 = v170 + v183
							}
						}
						v198 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v198 + v68
						*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v190
						return
					}
				}
			} else {
				v108 = v75
				if v68 <= int32(0) {
					v190 = v108
				} else {
					v112 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsAdd[2]))
					v114 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsAdd[1]))
					if v68 != int32(1) {
						v121 = l0
						v124 = v108
						v131 = int32(0)
						for {
							v132 = int32(2)
							*(*int32)(unsafe.Add(mBase, uint32(v114+v124<<(uint(v132)%32)))) = v121
							v137 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v112+v124))) = uint8(v137)
							v140 = v124 + v137
							v144 = int32(3)
							v146 = v121 + v137
							if base.Ui32(v146) <= base.Ui32(v144) {
								v149 = v144
							} else {
								v149 = v146
							}
							*(*int32)(unsafe.Add(mBase, uint32(v114+v140<<(uint(v132)%32)))) = v149
							v152 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v112+v140))) = uint8(v152)
							v154 = int32(3)
							v156 = v149 + v152
							if base.Ui32(v156) <= base.Ui32(v154) {
								v159 = v154
							} else {
								v159 = v156
							}
							v160 = int32(2)
							v161 = v124 + v160
							v163 = v131 + v160
							if v163 != v68&int32(2147483646) {
								v121 = v159
								v124 = v161
								v131 = v163
								continue
							} else {
								break
							}
							break
						}
						if v68&int32(1) == int32(0) {
							v190 = v161
						} else {
							v167 = v159
							v170 = v161
							*(*int32)(unsafe.Add(mBase, uint32(v114+v170<<(uint(int32(2))%32)))) = v167
							v183 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v112+v170))) = uint8(v183)
							v190 = v170 + v183
						}
					} else {
						v167 = l0
						v170 = v108
						*(*int32)(unsafe.Add(mBase, uint32(v114+v170<<(uint(int32(2))%32)))) = v167
						v183 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v112+v170))) = uint8(v183)
						v190 = v170 + v183
					}
				}
				v198 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v198 + v68
				*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v190
				return
			}
		}
	} else {
		v98 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
		if v98 < v75+v68 {
			F_KnownAssignedXidsCompress(m, int32(0), l2)
			mBase = m.M
			v103 = m.ExcPending
			if v103 != 0 {
				return
			} else {
				v104 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
				v105 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
				if v104 < v105+v68 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v221 = m.ExcPending
					if v221 != 0 {
						return
					} else {
						F_errmsg_internal(m, int32(_a_F_KnownAssignedXidsAdd_4), int32(0))
						mBase = m.M
						v225 = m.ExcPending
						if v225 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_KnownAssignedXidsAdd_1), int32(_a_F_KnownAssignedXidsAdd_5), int32(_a_F_KnownAssignedXidsAdd_3))
							mBase = m.M
							v230 = m.ExcPending
							if v230 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v108 = v105
					if v68 <= int32(0) {
						v190 = v108
					} else {
						v112 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsAdd[2]))
						v114 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsAdd[1]))
						if v68 != int32(1) {
							v121 = l0
							v124 = v108
							v131 = int32(0)
							for {
								v132 = int32(2)
								*(*int32)(unsafe.Add(mBase, uint32(v114+v124<<(uint(v132)%32)))) = v121
								v137 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v112+v124))) = uint8(v137)
								v140 = v124 + v137
								v144 = int32(3)
								v146 = v121 + v137
								if base.Ui32(v146) <= base.Ui32(v144) {
									v149 = v144
								} else {
									v149 = v146
								}
								*(*int32)(unsafe.Add(mBase, uint32(v114+v140<<(uint(v132)%32)))) = v149
								v152 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v112+v140))) = uint8(v152)
								v154 = int32(3)
								v156 = v149 + v152
								if base.Ui32(v156) <= base.Ui32(v154) {
									v159 = v154
								} else {
									v159 = v156
								}
								v160 = int32(2)
								v161 = v124 + v160
								v163 = v131 + v160
								if v163 != v68&int32(2147483646) {
									v121 = v159
									v124 = v161
									v131 = v163
									continue
								} else {
									break
								}
								break
							}
							if v68&int32(1) == int32(0) {
								v190 = v161
							} else {
								v167 = v159
								v170 = v161
								*(*int32)(unsafe.Add(mBase, uint32(v114+v170<<(uint(int32(2))%32)))) = v167
								v183 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v112+v170))) = uint8(v183)
								v190 = v170 + v183
							}
						} else {
							v167 = l0
							v170 = v108
							*(*int32)(unsafe.Add(mBase, uint32(v114+v170<<(uint(int32(2))%32)))) = v167
							v183 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v112+v170))) = uint8(v183)
							v190 = v170 + v183
						}
					}
					v198 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v198 + v68
					*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v190
					return
				}
			}
		} else {
			v108 = v75
			if v68 <= int32(0) {
				v190 = v108
			} else {
				v112 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsAdd[2]))
				v114 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsAdd[1]))
				if v68 != int32(1) {
					v121 = l0
					v124 = v108
					v131 = int32(0)
					for {
						v132 = int32(2)
						*(*int32)(unsafe.Add(mBase, uint32(v114+v124<<(uint(v132)%32)))) = v121
						v137 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v112+v124))) = uint8(v137)
						v140 = v124 + v137
						v144 = int32(3)
						v146 = v121 + v137
						if base.Ui32(v146) <= base.Ui32(v144) {
							v149 = v144
						} else {
							v149 = v146
						}
						*(*int32)(unsafe.Add(mBase, uint32(v114+v140<<(uint(v132)%32)))) = v149
						v152 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v112+v140))) = uint8(v152)
						v154 = int32(3)
						v156 = v149 + v152
						if base.Ui32(v156) <= base.Ui32(v154) {
							v159 = v154
						} else {
							v159 = v156
						}
						v160 = int32(2)
						v161 = v124 + v160
						v163 = v131 + v160
						if v163 != v68&int32(2147483646) {
							v121 = v159
							v124 = v161
							v131 = v163
							continue
						} else {
							break
						}
						break
					}
					if v68&int32(1) == int32(0) {
						v190 = v161
					} else {
						v167 = v159
						v170 = v161
						*(*int32)(unsafe.Add(mBase, uint32(v114+v170<<(uint(int32(2))%32)))) = v167
						v183 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v112+v170))) = uint8(v183)
						v190 = v170 + v183
					}
				} else {
					v167 = l0
					v170 = v108
					*(*int32)(unsafe.Add(mBase, uint32(v114+v170<<(uint(int32(2))%32)))) = v167
					v183 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v112+v170))) = uint8(v183)
					v190 = v170 + v183
				}
			}
			v198 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
			*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v198 + v68
			*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v190
			return
		}
	}
}
func F_kill(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	if l0 == int32(42) {
		v5 = F_raise(m, l1)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return int32(0)
		} else {
			return v5
		}
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_kill[0])) = int32(63)
		return int32(-1)
	}
}
