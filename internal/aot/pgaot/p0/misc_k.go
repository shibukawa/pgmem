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
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	v12 = *(*int32)(unsafe.Add(mBase, _consts[30]))
	if base.Ui32(l1) < base.Ui32(l0) {
		v14 = int32(1)
		if base.B2i32(base.Ui32(int32(2)) < base.Ui32(l1))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(l0)) == int32(0) {
			v26 = base.B2i32(base.Ui32(l0) < base.Ui32(l1))
		} else {
			v26 = int32(base.Ui32(l0-l1) >> (uint(int32(31)) % 32))
		}
		if v26 == int32(0) {
			v66 = v14
		} else {
			v32 = l0
			v33 = v14
			for {
				v39 = int32(1)
				v40 = v33 + v39
				v41 = int32(3)
				v43 = v32 + v39
				if base.Ui32(v43) <= base.Ui32(v41) {
					v46 = v41
				} else {
					v46 = v43
				}
				if base.B2i32(base.Ui32(int32(2)) < base.Ui32(l1))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v46)) == int32(0) {
					v58 = base.B2i32(base.Ui32(v46) < base.Ui32(l1))
				} else {
					v58 = int32(base.Ui32(v46-l1) >> (uint(int32(31)) % 32))
				}
				if v58 != 0 {
					v32 = v46
					v33 = v40
					continue
				} else {
					break
				}
				break
			}
			v66 = v40
		}
	} else {
		v66 = l1 - l0 + int32(1)
	}
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	if v73 < v72 {
		v76 = *(*int32)(unsafe.Add(mBase, _consts[604]))
		v77 = int32(2)
		v82 = *(*int32)(unsafe.Add(mBase, uint32(v76+v72<<(uint(v77)%32)-int32(4))))
		if base.B2i32(base.Ui32(v77) < base.Ui32(l0))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v82)) == int32(0) {
			v94 = base.B2i32(base.Ui32(l0) <= base.Ui32(v82))
		} else {
			v94 = base.B2i32(int32(0) <= v82-l0)
		}
		if v94 != 0 {
			F_KnownAssignedXidsDisplay(m, int32(15))
			mBase = m.M
			v198 = m.ExcPending
			if v198 != 0 {
				return
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v202 = m.ExcPending
				if v202 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(170657), int32(0))
					mBase = m.M
					v206 = m.ExcPending
					if v206 != 0 {
						return
					} else {
						F_errfinish(m, int32(483637), int32(4831), int32(455525))
						mBase = m.M
						v211 = m.ExcPending
						if v211 != 0 {
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
			v95 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
			if v95 < v72+v66 {
				F_KnownAssignedXidsCompress(m, int32(0), l2)
				mBase = m.M
				v100 = m.ExcPending
				if v100 != 0 {
					return
				} else {
					v101 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
					v102 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
					if v101 < v102+v66 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v215 = m.ExcPending
						if v215 != 0 {
							return
						} else {
							F_errmsg_internal(m, int32(170630), int32(0))
							mBase = m.M
							v219 = m.ExcPending
							if v219 != 0 {
								return
							} else {
								F_errfinish(m, int32(483637), int32(4848), int32(455525))
								mBase = m.M
								v224 = m.ExcPending
								if v224 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v105 = v102
						if v66 <= int32(0) {
							v185 = v105
						} else {
							v109 = *(*int32)(unsafe.Add(mBase, _consts[605]))
							v111 = *(*int32)(unsafe.Add(mBase, _consts[604]))
							if v66 != int32(1) {
								v116 = l0
								v119 = v105
								v123 = int32(0)
								for {
									v126 = int32(2)
									*(*int32)(unsafe.Add(mBase, uint32(v111+v119<<(uint(v126)%32)))) = v116
									v131 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v109+v119))) = uint8(v131)
									v134 = v119 + v131
									v138 = int32(3)
									v140 = v116 + v131
									if base.Ui32(v140) <= base.Ui32(v138) {
										v143 = v138
									} else {
										v143 = v140
									}
									*(*int32)(unsafe.Add(mBase, uint32(v111+v134<<(uint(v126)%32)))) = v143
									v146 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v109+v134))) = uint8(v146)
									v148 = int32(3)
									v150 = v143 + v146
									if base.Ui32(v150) <= base.Ui32(v148) {
										v153 = v148
									} else {
										v153 = v150
									}
									v154 = int32(2)
									v155 = v119 + v154
									v157 = v123 + v154
									if v157 != v66&int32(2147483646) {
										v116 = v153
										v119 = v155
										v123 = v157
										continue
									} else {
										break
									}
									break
								}
								v159 = v153
								v162 = v155
							} else {
								v159 = l0
								v162 = v105
							}
							if v66&int32(1) == int32(0) {
								v185 = v162
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v111+v162<<(uint(int32(2))%32)))) = v159
								v178 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v109+v162))) = uint8(v178)
								v185 = v162 + v178
							}
						}
						v192 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v192 + v66
						*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v185
						return
					}
				}
			} else {
				v105 = v72
				if v66 <= int32(0) {
					v185 = v105
				} else {
					v109 = *(*int32)(unsafe.Add(mBase, _consts[605]))
					v111 = *(*int32)(unsafe.Add(mBase, _consts[604]))
					if v66 != int32(1) {
						v116 = l0
						v119 = v105
						v123 = int32(0)
						for {
							v126 = int32(2)
							*(*int32)(unsafe.Add(mBase, uint32(v111+v119<<(uint(v126)%32)))) = v116
							v131 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v109+v119))) = uint8(v131)
							v134 = v119 + v131
							v138 = int32(3)
							v140 = v116 + v131
							if base.Ui32(v140) <= base.Ui32(v138) {
								v143 = v138
							} else {
								v143 = v140
							}
							*(*int32)(unsafe.Add(mBase, uint32(v111+v134<<(uint(v126)%32)))) = v143
							v146 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v109+v134))) = uint8(v146)
							v148 = int32(3)
							v150 = v143 + v146
							if base.Ui32(v150) <= base.Ui32(v148) {
								v153 = v148
							} else {
								v153 = v150
							}
							v154 = int32(2)
							v155 = v119 + v154
							v157 = v123 + v154
							if v157 != v66&int32(2147483646) {
								v116 = v153
								v119 = v155
								v123 = v157
								continue
							} else {
								break
							}
							break
						}
						v159 = v153
						v162 = v155
					} else {
						v159 = l0
						v162 = v105
					}
					if v66&int32(1) == int32(0) {
						v185 = v162
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v111+v162<<(uint(int32(2))%32)))) = v159
						v178 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v109+v162))) = uint8(v178)
						v185 = v162 + v178
					}
				}
				v192 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v192 + v66
				*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v185
				return
			}
		}
	} else {
		v95 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
		if v95 < v72+v66 {
			F_KnownAssignedXidsCompress(m, int32(0), l2)
			mBase = m.M
			v100 = m.ExcPending
			if v100 != 0 {
				return
			} else {
				v101 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
				v102 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
				if v101 < v102+v66 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v215 = m.ExcPending
					if v215 != 0 {
						return
					} else {
						F_errmsg_internal(m, int32(170630), int32(0))
						mBase = m.M
						v219 = m.ExcPending
						if v219 != 0 {
							return
						} else {
							F_errfinish(m, int32(483637), int32(4848), int32(455525))
							mBase = m.M
							v224 = m.ExcPending
							if v224 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v105 = v102
					if v66 <= int32(0) {
						v185 = v105
					} else {
						v109 = *(*int32)(unsafe.Add(mBase, _consts[605]))
						v111 = *(*int32)(unsafe.Add(mBase, _consts[604]))
						if v66 != int32(1) {
							v116 = l0
							v119 = v105
							v123 = int32(0)
							for {
								v126 = int32(2)
								*(*int32)(unsafe.Add(mBase, uint32(v111+v119<<(uint(v126)%32)))) = v116
								v131 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v109+v119))) = uint8(v131)
								v134 = v119 + v131
								v138 = int32(3)
								v140 = v116 + v131
								if base.Ui32(v140) <= base.Ui32(v138) {
									v143 = v138
								} else {
									v143 = v140
								}
								*(*int32)(unsafe.Add(mBase, uint32(v111+v134<<(uint(v126)%32)))) = v143
								v146 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v109+v134))) = uint8(v146)
								v148 = int32(3)
								v150 = v143 + v146
								if base.Ui32(v150) <= base.Ui32(v148) {
									v153 = v148
								} else {
									v153 = v150
								}
								v154 = int32(2)
								v155 = v119 + v154
								v157 = v123 + v154
								if v157 != v66&int32(2147483646) {
									v116 = v153
									v119 = v155
									v123 = v157
									continue
								} else {
									break
								}
								break
							}
							v159 = v153
							v162 = v155
						} else {
							v159 = l0
							v162 = v105
						}
						if v66&int32(1) == int32(0) {
							v185 = v162
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v111+v162<<(uint(int32(2))%32)))) = v159
							v178 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v109+v162))) = uint8(v178)
							v185 = v162 + v178
						}
					}
					v192 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v192 + v66
					*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v185
					return
				}
			}
		} else {
			v105 = v72
			if v66 <= int32(0) {
				v185 = v105
			} else {
				v109 = *(*int32)(unsafe.Add(mBase, _consts[605]))
				v111 = *(*int32)(unsafe.Add(mBase, _consts[604]))
				if v66 != int32(1) {
					v116 = l0
					v119 = v105
					v123 = int32(0)
					for {
						v126 = int32(2)
						*(*int32)(unsafe.Add(mBase, uint32(v111+v119<<(uint(v126)%32)))) = v116
						v131 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v109+v119))) = uint8(v131)
						v134 = v119 + v131
						v138 = int32(3)
						v140 = v116 + v131
						if base.Ui32(v140) <= base.Ui32(v138) {
							v143 = v138
						} else {
							v143 = v140
						}
						*(*int32)(unsafe.Add(mBase, uint32(v111+v134<<(uint(v126)%32)))) = v143
						v146 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v109+v134))) = uint8(v146)
						v148 = int32(3)
						v150 = v143 + v146
						if base.Ui32(v150) <= base.Ui32(v148) {
							v153 = v148
						} else {
							v153 = v150
						}
						v154 = int32(2)
						v155 = v119 + v154
						v157 = v123 + v154
						if v157 != v66&int32(2147483646) {
							v116 = v153
							v119 = v155
							v123 = v157
							continue
						} else {
							break
						}
						break
					}
					v159 = v153
					v162 = v155
				} else {
					v159 = l0
					v162 = v105
				}
				if v66&int32(1) == int32(0) {
					v185 = v162
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v111+v162<<(uint(int32(2))%32)))) = v159
					v178 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v109+v162))) = uint8(v178)
					v185 = v162 + v178
				}
			}
			v192 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
			*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v192 + v66
			*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v185
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
		*(*int32)(unsafe.Add(mBase, _consts[166])) = int32(63)
		return int32(-1)
	}
}
